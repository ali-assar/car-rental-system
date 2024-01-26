# Aggregator Microservice

The Aggregator microservice receives distance data from the Distance Calculator, aggregates the data, and provides endpoints for retrieving aggregated information. The service is instrumented with Prometheus for monitoring.

## Overview

The Aggregator microservice is designed to handle both gRPC and HTTP requests. It aggregates distance data and exposes endpoints to retrieve aggregated information.

## Prerequisites

Before running the Aggregator microservice, ensure you have the following prerequisites installed:

- Go programming language
- Prometheus for monitoring (optional but recommended)
- Optionally, a store type (e.g., "memory") configured in the `.env` file

## Configuration
The Aggregator microservice uses environment variables for configuration. Modify the .env file with the following details:

```
AGG_GRPC_ENDPOINT=your-gRPC-endpoint
AGG_HTTP_ENDPOINT=your-HTTP-endpoint
AGG_STORE_TYPE=memory  # or choose another valid store type
```

## Endpoints

`/aggregate`
This endpoint returns aggregated distance data.

`/invoice`
This endpoint returns invoice information.

`/metrics`
Prometheus metrics endpoint for monitoring the Aggregator microservice.

## RUN aggregator
you can easily run this service by 
```
make agg
```
note to have data of dependent services you need to also run
```
make calculator
make receiver
make obu
```