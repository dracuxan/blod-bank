.PHONY: server client proto

server:
	go run server/main.go

client:
	@mkdir -p client
	@go build -o bin/client client/main.go

test:
	go test ./... -v

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/blod.proto
