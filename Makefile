APP_NAME=go-gin-todo

.PHONY: init lint test build run docker-build
init:
	go mod tidy

lint:
	golangci-lint run

test:
	go test ./...

build:
	go build -o bin/$(APP_NAME) ./cmd/server

run:
	go run ./cmd/server

docker-build:
	docker build -t $(APP_NAME):local .
