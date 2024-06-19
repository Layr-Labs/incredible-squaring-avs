// Package httpd provides the HTTP server for accessing the distributed key-value store.
// It also provides the endpoint for other nodes to join an existing cluster.
package operator

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/ethereum/go-ethereum/crypto"
)

type IPriceFSM interface {
	Join(nodeID, addr string) error
	TriggerElection()
}

type Service struct {
	addr                  string
	ln                    net.Listener
	taskSubmissions       map[string]int
	logger                logging.Logger
	priceFSM              IPriceFSM
	taskResponsesMu       sync.RWMutex
	taskResponses         *map[uint32]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse
	blsAggregationService blsagg.BlsAggregationService
	avsReader             chainio.AvsReaderer
	ethClient             eth.Client
}

// New returns an uninitialized HTTP service.
func NewService(addr string, priceFSM IPriceFSM, blsAggregationService blsagg.BlsAggregationService, taskResponses *map[uint32]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse, ethClient eth.Client) *Service {
	return &Service{
		addr:                  addr,
		priceFSM:              priceFSM,
		taskResponses:         taskResponses,
		blsAggregationService: blsAggregationService,
		ethClient:             ethClient,
	}
}

// Start starts the service.
func (s *Service) Start() error {
	server := http.Server{
		Handler: s,
	}

	log.Println("Initializing http server")

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.ln = ln

	http.Handle("/", s)

	go func() {
		err := server.Serve(s.ln)
		if err != nil {
			log.Fatalf("HTTP serve: %s", err)
		}
	}()

	log.Printf("Http server started\n")

	return nil
}

// ServeHTTP allows Service to serve HTTP requests.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/submitAvsTask" {
		s.handlePriceUpdateTaskSubmittion(w, r)
	} else if r.URL.Path == "/join" {
		s.handleJoin(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) handleJoin(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(m) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	signedMessage, ok := m["signedMessage"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageHash, ok := m["messageHash"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blockNumber, ok := m["blockNumber"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	signedMessageBytes, err := base64.StdEncoding.DecodeString(signedMessage)

	if err != nil {
		s.logger.Warn("Failed to decode signed message", "err", err)
	}

	messageBytes, err := base64.StdEncoding.DecodeString(messageHash)

	if err != nil {
		s.logger.Warn("Failed to decode message hash", "err", err)
	}

	sigPublicKey, err := crypto.SigToPub(messageBytes, signedMessageBytes)

	if err != nil {
		s.logger.Warn("Failed to parse operator signature", "err", err)
	}

	validOperatorUrls, err := s.avsReader.FetchOperatorUrl(context.Background(), crypto.PubkeyToAddress(*sigPublicKey))

	if err != nil {
		s.logger.Warn("Resolved address is not a valid operator", "address", crypto.PubkeyToAddress(*sigPublicKey), "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Resolved address is not a valid operator"))
		return
	} else {
		s.logger.Warn("Resolved operator address joining raft cluster is", "address", crypto.PubkeyToAddress(*sigPublicKey))
	}

	data := []byte(blockNumber)
	hash := crypto.Keccak256Hash(data)
	resolvedBlockNumberHash := base64.StdEncoding.EncodeToString(hash.Bytes()[:])

	if messageHash != resolvedBlockNumberHash {
		s.logger.Warn("Blocknumber hash does not match block number")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Blocknumber hash does not match block number"))
		return
	}

	// Verify block number is within last 2 blocks to protect against stale signatures
	latestBlock, _ := s.ethClient.BlockNumber(context.Background())

	blockAsInt, _ := strconv.ParseUint(blockNumber, 10, 64)

	if blockAsInt != latestBlock && blockAsInt != latestBlock-1 {
		s.logger.Warn("Blocknumber in signature is to old")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Blocknumber in signature is to old"))
		return
	}

	nodeID := crypto.PubkeyToAddress(*sigPublicKey).String()

	if err := s.priceFSM.Join(nodeID, validOperatorUrls.RpcUrl); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) handlePriceUpdateTaskSubmittion(w http.ResponseWriter, r *http.Request) {

	var signedResponse SignedTaskResponse

	if err := json.NewDecoder(r.Body).Decode(&signedResponse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Submit each price feed source seperatly
	s.logger.Info("Preparing to submit bls signatures")
	for i, task := range signedResponse.TaskResponse {
		taskIndex := task.TaskId
		signature := signedResponse.BlsSignature[i]
		taskResponseDigest, err := core.GetTaskResponseDigest(task.Price, task.Source, task.TaskId, task.Decimals)
		if err != nil {
			s.logger.Error("Failed to get task response digest", "err", err)
			w.WriteHeader(http.StatusBadGateway)
		}
		s.taskResponsesMu.Lock()

		if _, ok := (*s.taskResponses)[taskIndex]; !ok {
			(*s.taskResponses)[taskIndex] = make(map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse)
		}
		if _, ok := (*s.taskResponses)[taskIndex][taskResponseDigest]; !ok {
			(*s.taskResponses)[taskIndex][taskResponseDigest] = cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse{
				Price:    uint32(task.Price),
				Decimals: 18,
				Source:   task.Source,
				TaskId:   task.TaskId,
			}
		}
		s.taskResponsesMu.Unlock()

		err = s.blsAggregationService.ProcessNewSignature(
			context.Background(), taskIndex, taskResponseDigest,
			&signature, signedResponse.OperatorId,
		)

		s.logger.Info("Submitted bls signature to aggregation service",
			"taskId", task.TaskId,
			"operatorId", signedResponse.OperatorId.LogValue().String(),
		)
	}
}
