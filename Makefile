build:
	@go build -o bin/catfact

run: build
	@ ./bin/catfact

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto

.PHONY: proto