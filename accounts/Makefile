SERVICE=accounts

default: lint

build:
	go build -v ./cmd/$(SERVICE)

lint:
	golangci-lint run -v ./...


create:
	protoc --proto_path=proto ./proto/*.proto --go_out=internal/api/
	protoc --proto_path=proto ./proto/*.proto --go-grpc_out=internal/api/

clian:
	rm internal/api/proto/*.go


.PHONY: build lint create clean/api/proto/*.go