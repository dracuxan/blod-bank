package main

import (
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	"github.com/dracuxan/blod-bank/client/runner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var configServer = "localhost:5000"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./client <command> [command-flags]")
		fmt.Println("Commands: get, register, delete, list, update, server")
		os.Exit(1)
	}

	option := os.Args[1]

	conn, err := grpc.NewClient(configServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	helper.CheckCommonError(err, "did not connect")

	defer conn.Close()

	runner.Run(conn, option)
}
