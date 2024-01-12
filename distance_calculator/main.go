package main

import "log"

var kafkaTopic = "obudata"

func main() {
	//var svc CalculatorService
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
