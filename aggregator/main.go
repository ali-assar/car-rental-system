package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	var (
		store          = makeStore()
		svc            = NewInvoiceAggregator(store)
		grpcListenAddr = os.Getenv("AGG_GRPC_ENDPOINT")
		httpListenAddr = os.Getenv("AGG_HTTP_ENDPOINT")
	)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)
	go func() {
		log.Fatal(makeGRPCTransport(grpcListenAddr, svc))
	}()
	log.Fatal(makeHTTPTransport(httpListenAddr, svc))
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
	aggMetricHandler := newHTTPMetricHandler("aggregate")
	invMetricHandler := newHTTPMetricHandler("invoice")
	http.HandleFunc("/aggregate", aggMetricHandler.instrument(handleAggregate(svc)))
	http.HandleFunc("/invoice", invMetricHandler.instrument(handleGetInvoice(svc)))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("HTTP transport running on port", listenAddr)
	return (http.ListenAndServe(listenAddr, nil))
}

func makeStore() Storer {
	storeType := os.Getenv("AGG_STORE_TYPE")
	switch storeType {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatalf("invalid store type given %s", storeType)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}
