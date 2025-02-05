.DEFAULT_GOAL := build

fmt: 
	@go fmt ./...

lint: fmt
	@golint ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o bin/main cmd/server/main.go

clean:
	@go clean
	rm bin/main

test:
	@go test ./...

.PHONY: fmt lint vet build clean
