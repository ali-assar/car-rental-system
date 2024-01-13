api:
	@go build -o bin/api
	@./bin/api

seed:
	@go run scripts/seed.go

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

.PHONY: deps build run seed obu receiver calculator agg test api
