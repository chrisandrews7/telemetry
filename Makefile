GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

include .env
export

tools: bin/mockery
bin/mockery:
	GOBIN=${GOBIN} go install github.com/vektra/mockery/v2/...

.PHONY: generate-mocks
generate-mocks: tools
	GOBIN=${GOBIN} GOBASE=${GOBASE} go generate ./...

.PHONY: test
test: 
	go test ./internal/...

.PHONY: integration-test
integration-test: 
	go test ./cmd/telemetry/app --count=1

.PHONY: generate-telemetry
generate-telemetry: 
	./build/darwin-amd64/telemetry generate

.PHONY: start-deps
start-deps:
	docker-compose up

.PHONY: start
start:
	go run ./cmd/telemetry/main.go