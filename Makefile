all: test lint vet

test: dependencies
	go test ./...

lint: dependencies
	@which golint || go get github.com/golang/lint/golint
	golint ./...

vet: dependencies
	go vet ./...

dependencies:
	go get -t

.PHONY: dependencies lint test
