// Package httpd provides the HTTP server for accessing the distributed key-value store.
// It also provides the endpoint for other nodes to join an existing cluster.
package operator

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
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
}

// New returns an uninitialized HTTP service.
func NewService(addr string, priceFSM IPriceFSM, blsAggregationService blsagg.BlsAggregationService, taskResponses *map[uint32]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse) *Service {
	return &Service{
		addr:                  addr,
		priceFSM:              priceFSM,
		taskResponses:         taskResponses,
		blsAggregationService: blsAggregationService,
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
	} else if r.URL.Path == "/verify" {
		s.handleVerify(w, r)
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

	if len(m) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	remoteAddr, ok := m["addr"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nodeID, ok := m["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 1 - Add call to send a message requesting it be signed by remoteAddr
	// 2 - Verify remote add is a valid operator by verifying signature

	if err := s.priceFSM.Join(nodeID, remoteAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

func (s *Service) handleVerify(w http.ResponseWriter, r *http.Request) {
	// 1 - Sign a message with nodeId + operator address
	// 2 - Return signed message
}
