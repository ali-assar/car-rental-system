package main

import (
	"net"
	"net/http"
	"os"

	"github.com/Ali-Assar/car-rental-system/go-kit-example/aggsvc/aggendpoint"
	"github.com/Ali-Assar/car-rental-system/go-kit-example/aggsvc/aggservice"
	"github.com/Ali-Assar/car-rental-system/go-kit-example/aggsvc/aggtransport"
	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := aggservice.New(logger)
	endpoints := aggendpoint.New(service, logger)
	httpHandler := aggtransport.NewHTTPHandler(endpoints, logger)
	//	grpcServer     = addtransport.NewGRPCServer(endpoints, tracer, zipkinTracer, logger)
	// The HTTP listener mounts the Go kit HTTP handler we created.
	httpListener, err := net.Listen("tcp", ":3000")
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	httpHandler = aggtransport.NewHTTPHandler(endpoints, logger)
	logger.Log("transport", "HTTP", "addr", ":3000")
	err = http.Serve(httpListener, httpHandler)
	if err != nil {
		panic(err)
	}

}
