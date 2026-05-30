.PHONY: fmt validate build run

default: run

fmt:
	go fmt ./...

validate:
	go vet ./...

build:
	go build -o ./bin/web ./cmd/web

run:
	go run ./cmd/web

