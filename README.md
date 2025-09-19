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
│       └── run.go
├── proto
│   ├── blod_grpc.pb.go
│   ├── blod.pb.go
│   └── blod.proto
├── flake.lock
├── flake.nix
├── go.mod
├── go.sum
├── README.md
└── LICENSE
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
