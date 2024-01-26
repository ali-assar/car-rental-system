# Distance Calculator Microservice

The Distance Calculator microservice consumes OBU data from a Kafka topic, calculates the distance based on latitude and longitude, and forwards the processed data to the Aggregator service.

## Overview

The Distance Calculator microservice is designed to receive OBU data from a Kafka topic, calculate the distance, and communicate with the Aggregator service either via gRPC or HTTP. The calculated data is then sent to the Aggregator for further processing.

## Prerequisites

Before running the Distance Calculator microservice, ensure you have the following prerequisites installed:

- Go programming language
- Kafka broker running and accessible
- Aggregator service running and reachable at the specified endpoint (HTTP or gRPC)

## RUN distance-calculator
you can easily run this service by 
```
make calculator
```
note to consume data you need to also run
```
make receiver
make obu
```
