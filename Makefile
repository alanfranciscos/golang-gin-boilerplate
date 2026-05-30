BINARY_NAME=api

.PHONY: build test format lint docker-build docker-test

build:
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

test:
	go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out | awk '$$3 ~ /^[0-9.]+%/ { if (substr($$3, 1, length($$3)-1) < 90) { print "Coverage for " $$1 " is below 90%: " $$3; exit 1 } }'

format:
	go fmt ./...

lint:
	@echo "Running lint..."
	@# Assuming golangci-lint is installed
	golangci-lint run

docker-build:
	docker build -t go-boilerplate .

docker-test:
	docker build --target test -t go-boilerplate-test .
