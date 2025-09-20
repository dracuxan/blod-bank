# blod-bank

A gRPC-based configuration management service with a simple CLI client.
Currently supports fetching and listing stored YAML-like config files.

## Project Structure

```sh
blod-bank
│
├── server
│   └── main.go
├── client
│   ├── helper
│   │   └── helper.go
│   ├── main.go
│   └── runner
│       ├── commands
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── list.go
│       │   ├── register.go
│       │   └── update.go
│       └── run.go
├── proto
│   ├── blod_grpc.pb.go
│   ├── blod.pb.go
│   └── blod.proto
├── flake.lock
├── flake.nix
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

## Getting Started

1. Clone the repo
2. Run the server:
   ```sh
   go run server/main.go
   ```
3. Run the client:
   ```sh
   go run client/main.go
   ```

## Roadmap

- [x] Get single config
- [x] List all config
- [x] Register new config
- [x] Update config
- [x] Delete config
- [x] Basic CLI tool

## TODO

- [ ] Add database connection
- [ ] Replace in-memory storage with Postgress in server
