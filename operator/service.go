// Package httpd provides the HTTP server for accessing the distributed key-value store.
// It also provides the endpoint for other nodes to join an existing cluster.
package operator

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type IPriceFSM interface {
	Join(nodeID, addr string) error
	TriggerElection()
}

type Service struct {
	addr            string
	ln              net.Listener
	taskSubmissions map[string]int

	priceFSM IPriceFSM
}

// New returns an uninitialized HTTP service.
func NewService(addr string, priceFSM IPriceFSM) *Service {
	return &Service{
		addr:     addr,
		priceFSM: priceFSM,
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

	m := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Resolved body is %+v\n", m)

	// First example just log response
	// 1 - Verify senders bs signature + ensure operator processing is leader
	// 2 - Verify price submissions fall within error range
	// 3 - If this task has reached qurom submit it on-chain and trigger a new leader election
}

func (s *Service) handleVerify(w http.ResponseWriter, r *http.Request) {
	// 1 - Sign a message with nodeId + operator address
	// 2 - Return signed message
}
