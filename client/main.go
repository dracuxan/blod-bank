package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// statusAddr       = flag.String("statusAddr", "localhost:5000", "server address to check server status")
var configServer = flag.String("configServerAddr", "localhost:5001", "server address for Donor Service")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*configServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
}
