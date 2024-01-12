build:
	@go build -o bin/api
run: build
	@./bin/api
seed:
	@go run scripts/seed.go
obu:
	@go build -o bin/obu ./obu
	@./bin/obu
receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver
.PHONY: obu
calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

docker:
	@echo "building Docker image"
	@docker build -t api .
	@echo "running API inside Docker container"
	@docker run -p 3000:3000 api
test:
	@go test -count=1 -v ./...