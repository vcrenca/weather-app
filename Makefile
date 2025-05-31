.PHONY: help build test lint run test-coverage compose-up compose-down logs

help:
	@echo "Usage:"
	@echo "  make build         	Build the Go binary"
	@echo "  make test          	Run tests"
	@echo "  make test-coverage 	Run tests and open coverage report"
	@echo "  make lint          	Run linter (go vet)"
	@echo "  make run           	Run the app on host"
	@echo "  make compose-up    	Run the app in docker-compose"
	@echo "  make compose-down  	Stop the app in docker-compose"
	@echo "  make logs		Show the logs of docker-compose"

run-local:
	ENV=LOCAL go run cmd/main.go

build:
	go build -o bin/weather-app cmd/main.go

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

generate:
	go tool oapi-codegen -package v1 -generate "chi-server" -o internal/ports/http/v1/opeanpi_v1_server.gen.go api/v1/weather.yaml
	go tool oapi-codegen -package v1 -generate "types" -o internal/ports/http/v1/opeanpi_v1_types.gen.go api/v1/weather.yaml

lint:
	go vet ./...

compose-up:

	docker compose up --build -d

compose-down:
	docker compose down

logs:
	docker compose logs -f
