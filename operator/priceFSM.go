package operator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	priceFeedAdapter "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/PriceFeedAdapter"
	"github.com/Layr-Labs/incredible-squaring-avs/core"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

const (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

type priceUpdateCommand struct {
	FeedName string `json:"feedName"`
}

type SignedTaskResponse struct {
	TaskResponse []PriceUpdateTaskResponse
	BlsSignature []bls.Signature
	OperatorId   sdktypes.OperatorId
}

type PriceFSM struct {
	RaftDir      string // Directory for operator raft logs
	RaftBind     string // host:port used by the operator for raft protocol
	RaftHttpBind string // host:port for custom server for custom raft logic
	raft         *raft.Raft
	mu           sync.Mutex
	priceData    map[string]int // past price data
	logger       *log.Logger
	// needed to fetch the price of assets on different on-chain oracle networks
	priceFeedAdapter *priceFeedAdapter.ContractPriceFeedAdapter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
}

func NewConcensusFSM(feedAdapter *priceFeedAdapter.ContractPriceFeedAdapter, keyPair *bls.KeyPair) *PriceFSM {
	return &PriceFSM{
		priceData:        make(map[string]int),
		logger:           log.New(os.Stderr, "[priceData] ", log.LstdFlags),
		priceFeedAdapter: feedAdapter,
		blsKeypair:       keyPair,
	}
}

func JoinExistingNetwork(joinAddr, raftAddr, nodeID string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr, "id": nodeID})
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/join", joinAddr), "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	log.Printf("Joined raft consensus through uri %s:", joinAddr)
	defer resp.Body.Close()
	return nil
}

func (p *PriceFSM) setOperatorId(id sdktypes.OperatorId) {
	p.operatorId = id
}

// Operator initializes raft consenses server.
// If enableSingle is set, and there are no existing peers,
// then this node becomes the first node, and therefore leader, of the cluster.
// localID should be the server identifier for this node.
func (p *PriceFSM) Initialize(enableSingle bool, localId string) error {
	// Setup Raft configuration.
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localId)

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", p.RaftBind)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(p.RaftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(p.RaftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}

	// Create the log store and stable store using BoltDB in memory key value store
	var logStore raft.LogStore
	var stableStore raft.StableStore

	boltDB, err := raftboltdb.New(raftboltdb.Options{
		Path: filepath.Join(p.RaftDir, "raft.db"),
	})
	if err != nil {
		return fmt.Errorf("new bbolt store: %s", err)
	}

	logStore = boltDB
	stableStore = boltDB

	// Instantiate the Raft systems.
	ra, err := raft.NewRaft(config, (*fsm)(p), logStore, stableStore, snapshots, transport)
	if err != nil {
		return fmt.Errorf("new raft: %s", err)
	}
	p.raft = ra

	// If only node and not joining an existing raft network bootstrap the network
	if enableSingle {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	}
	return nil
}

// Join joins a node, identified by nodeID and located at addr, to this store.
// The node must be ready to respond to Raft communications at that address.
func (p *PriceFSM) Join(nodeID, addr string) error {
	p.logger.Printf("received join request for remote node %s at %s", nodeID, addr)

	configFuture := p.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		p.logger.Printf("failed to get raft configuration: %v", err)
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				p.logger.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return nil
			}

			future := p.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := p.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	p.logger.Printf("node %s at %s joined successfully", nodeID, addr)
	return nil
}

func (p *PriceFSM) IsLeader() (bool, string) {
	leaderURL, _ := p.raft.LeaderWithID()
	return string(leaderURL) == p.RaftBind, string(leaderURL)
}

func (p *PriceFSM) TriggerElection() {
	p.raft.LeadershipTransfer()
}

type fsm PriceFSM

type fsmSnapshot struct {
	priceData map[string]int
}

// Triggers operator to fetch the requested price feed and sumbit to leader
func (f *fsm) Apply(l *raft.Log) interface{} {

	// Leader does not respond to task request from themselves
	leaderURL, _ := f.raft.LeaderWithID()

	if string(leaderURL) == f.RaftBind {
		return nil
	}

	lastAppliedIndex := f.raft.AppliedIndex()

	if l.Index < lastAppliedIndex {
		return nil // No need to replay previous logs
	}

	var request PriceUpdateRequest
	if err := json.Unmarshal(l.Data, &request); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}

	// Fetch chainlink price
	resolvePrice, err := f.priceFeedAdapter.GetLatestPrice(&bind.CallOpts{}, request.FeedName)

	if err != nil {
		f.logger.Printf("Failed to fetch price", "err", err)
		return nil
	}

	response := []PriceUpdateTaskResponse{} // slice will automatically resize if needed

	chainlinkResponse := PriceUpdateTaskResponse{Price: uint32(resolvePrice.Uint64()), Source: "chainlink", TaskId: request.TaskId, Decimals: 18}

	f.logger.Printf("Chainlink response: %v", chainlinkResponse)
	response = append(response, chainlinkResponse)

	if err := f.SubmitTaskToLeader(request, response, request.LeaderUrl); err != nil {
		f.logger.Printf("Failed to submit task response", "err", err)
	}

	return nil
}

func (f *fsm) SubmitTaskToLeader(request PriceUpdateRequest, responses []PriceUpdateTaskResponse, leaderUrl string) error {

	responseSignatures := []bls.Signature{}
	signedResponses := []PriceUpdateTaskResponse{}

	// Iterate over every response and sign via bls signature
	for i, response := range responses {
		if response.Source == "" {
			continue
		}

		f.logger.Printf("Submiting response %v for task %v\n", i, response.TaskId)
		taskResponseHash, err := core.GetTaskResponseDigest(response.Price, response.Source, response.TaskId, response.Decimals)
		if err != nil {
			log.Printf("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
			return err
		}
		responseSignatures = append(responseSignatures, *f.blsKeypair.SignMessage(taskResponseHash))
		signedResponses = append(signedResponses, response)
	}

	signedTaskResponse := SignedTaskResponse{
		TaskResponse: signedResponses,
		BlsSignature: responseSignatures,
		OperatorId:   f.operatorId,
	}

	b, err := json.Marshal(signedTaskResponse)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/submitAvsTask", leaderUrl), "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	log.Printf("Submitted task to %s:", leaderUrl)
	defer resp.Body.Close()
	return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	return &fsmSnapshot{priceData: make(map[string]int)}, nil
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	return nil
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := json.Marshal(f.priceData)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (f *fsmSnapshot) Release() {}
