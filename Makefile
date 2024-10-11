.PHONY: proto

chitty-run:
	@go run cmd/chitty/main.go

chitty-build:
	@go build -o bin/chitty cmd/chitty/main.go

chitty-install:
	@go install cmd/chitty/main.go

grpc-serve:
	@go run cmd/grpc/main.go

grpc-build:
	@go build -o bin/grpc cmd/grpc/main.go

proto: 
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/*.proto
