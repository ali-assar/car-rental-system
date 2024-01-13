package main

import (
	"log"

	"github.com/Ali-Assar/car-rental-system/aggregator/client"
)

const (
	kafkaTopic        = "obudata"
	aggregateEndPoint = "http://localhost:3000/aggregate"
)

func main() {
	//var svc CalculatorService
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregateEndPoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
