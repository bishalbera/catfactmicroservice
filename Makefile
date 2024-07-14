build:
	@go build -o bin/catfact

run: build
	@ ./bin/catfact

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto

install-express:
	@cd auth-service && npm install

run-express:
	@cd auth-service/src/ && npm start

.PHONY: proto