# blod-bank

A gRPC-based configuration management service with a simple CLI client.  
Currently supports fetching and listing stored YAML-like config files.

## Project Structure

```sh
blod-bank
│
├── server
│   └── main.go
├── status_server
│   └── main.go
├── client
│   └── main.go
├── proto
│   ├── blod_grpc.pb.go
│   ├── blod.pb.go
│   └── blod.proto
├── go.mod
├── go.sum
├── flake.lock
├── flake.nix
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
   ```
   go run client/main.go
   ```

## Roadmap

- [x] Get single config
- [x] List all config
- [ ] Register new config
- [ ] Update config
- [ ] Delete config
- [ ] Basic CLI tool
