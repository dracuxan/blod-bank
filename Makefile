.PHONY: all run build test clean help

run:
	@go run main.go

build:
	@go build -o bin/cli main.go

test:
	@go test -v ./...

clean:
	@rm -rf bin

help:
	@echo "Available commands:"
	@echo "  make run    - Run the application"
	@echo "  make build  - Build the application"
	@echo "  make test   - Run tests"
	@echo "  make clean  - Clean up build artifacts"
	@echo "  make help   - Show this help message"
