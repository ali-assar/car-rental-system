package main

import "log"

var kafkaTopic = "obudata"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
