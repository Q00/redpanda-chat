.PHONY: run vendor build

ROOT = $(PWD)

# Install go modules dependencies
vendor:
	go mod vendor
run:
	@go run ./cmd/chatd/main.go

test:
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

up:
	@docker-compose -f docker-compose.yml up

down:
	@docker-compose -f docker-compose..yml down -v

build:
	@CGO_ENABLED=0  go build -o ./engine ./cmd/chatd/main.go
