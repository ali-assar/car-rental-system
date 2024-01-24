api:
	@go build -o bin/api
	@./bin/api

seed:
	@go run scripts/seed.go

gate:
	@go build -o bin/gate ./gateway
	@./bin/gate
obu:
	@go build -o bin/obu ./obu
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

agg:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

test:
	@go test -count=1 -v ./...

all:  calculator receiver obu


proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

.PHONY: deps build run seed obu receiver calculator agg test api gate
