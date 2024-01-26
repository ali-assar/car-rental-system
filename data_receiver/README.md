# Data Receiver Microservice

The Data Receiver microservice is responsible for receiving simulated OBU (On-Board Unit) data and producing messages to a Kafka topic.

## Overview

The Data Receiver microservice is designed to handle WebSocket connections, receive OBU data, generate a unique RequestID, and produce the data to a Kafka topic for further processing by downstream services.

## RUN receiver

Start the Data Receiver microservice:

bash
Copy code
make receiver
make obu
