build:
	@go build -o bin/api

run: build
	@./bin/api

seed:
	@go run scripts/seed.go

docker:
	@echo "building Docker image"
	@docker build -t api .
	@echo "running API inside Docker container"
	@docker run -p 3000:3000 api

obu:
	@go run obu/main.go

.PHONY: obu

test:
	@go test -count=1 -v ./...
