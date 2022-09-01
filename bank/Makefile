create:
	protoc --proto_path=proto proto/*.proto --go_out=internal/api/gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=internal/api/gen/

clian:
	rm gen/proto/*.go