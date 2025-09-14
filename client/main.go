package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/blod-bank/client/runner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var configServer = "localhost:5001"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./client <command> --<optional sub commands>")
	}

	conn, err := grpc.NewClient(configServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	runner.Run(conn)
}
