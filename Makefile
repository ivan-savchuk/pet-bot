include .env

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	OPEN_WEATHER_API_KEY=$(OPEN_WEATHER_API_KEY) go run ./cmd/main/main.go