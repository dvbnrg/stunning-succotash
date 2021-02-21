.PHONY: run clean

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/userservice.proto

build: 
	@go build .

run: proto build
	./userService

clean: 
	@rm userService pb/userservice.pb.go pb/userservice_grpc.pb.go

