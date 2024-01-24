package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	httpListenAddr := flag.String("httpAddr", ":30000", "the listen address of HTTP server")
	grpcListenAddr := flag.String("grpcAddr", ":3001", "the listen address of GRPC server")
	flag.Parse()
	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)
	go func() {
		log.Fatal(makeGRPCTransport(*grpcListenAddr, svc))
	}()
	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)
	// Generate a TCP listener
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("stopping GRPC transport")
		listener.Close()
	}()
	// Make a new GRPC native server
	server := grpc.NewServer([]grpc.ServerOption{}...)
	//Register (our) GRPC server implementation to the GRPC package
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(listener)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("HTTP transport running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc))
	http.Handle("/metrics", promhttp.Handler())
	return (http.ListenAndServe(listenAddr, nil))
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		values, ok := r.URL.Query()["obu"]
		if !ok {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "OBU id is missing"})
			return
		}
		obuID, err := strconv.Atoi(values[0])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid OBU ID"})
			return
		}
		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
