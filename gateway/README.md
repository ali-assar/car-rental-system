# Gateway Microservice

The Gateway microservice serves as an HTTP interface to interact with the Aggregator service, providing an endpoint to retrieve invoice information.

## Overview

The Gateway microservice handles HTTP requests and communicates with the Aggregator service to retrieve invoice information.

## Prerequisites

Before running the Gateway microservice, ensure you have the following prerequisites installed:

- Go programming language
- Aggregator service running and accessible at the specified endpoint

## Getting Started

## Configuration
The Gateway microservice uses command-line flags for configuration. Modify the flags in the main.go file if needed:

```
-HTTP listenAddr=:6000
-aggServiceAddr=http://localhost:4000
```
## RUN gateway
you can easily run this service by 
```
make gate
```
note to have data of dependent services you need to also run
```
make agg
make calculator
make receiver
make obu
```