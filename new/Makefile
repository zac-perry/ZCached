.DEFAULT_GOAL := build

fmt: 
	@go fmt ./...

lint: fmt
	@golint ./...

vet: fmt
	@go vet ./...

build: vet
	@go build -o bin/cache . 

clean:
	@go clean
	rm bin/*

test:
	@go test ./...

run: build
	./bin/cache

.PHONY: fmt lint vet build clean
