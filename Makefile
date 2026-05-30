BINARY_NAME=api

.PHONY: build test format lint docker-build docker-test run docker-run

build:
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test ./...


format:
	go fmt ./...

lint:
	golangci-lint run

docker-build:
	docker build -t go-boilerplate .

docker-run:
	docker run -p 8080:8080 --env-file .env go-boilerplate

docker-test:
	docker build --target test -t go-boilerplate-test .
