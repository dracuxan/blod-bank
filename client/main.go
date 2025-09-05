package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var statusAddr = flag.String("statusAddr", "localhost:5000", "server address to connect")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*statusAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := blodBank.NewSystemServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &blodBank.NoParam{})
	if err != nil {
		log.Fatal("The server is DOWN")
	}
	fmt.Println(r.GetMessage())
}
