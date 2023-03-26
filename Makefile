include .env

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	AQI_API_KEY=$(AQI_API_KEY) go run ./cmd/main/main.go