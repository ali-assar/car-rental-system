package main

import (
	"log"

	"github.com/Ali-Assar/car-rental-system/aggregator/client"
)

const (
	kafkaTopic        = "obudata"
	aggregateEndPoint = "http://localhost:30000"
)

func main() {
	//var svc CalculatorService
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)

	httpClient := client.NewHTTPClient(aggregateEndPoint)
	//grpcClient, err := client.NewGRPCClient(aggregateEndPoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
