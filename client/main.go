package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// statusAddr       = flag.String("statusAddr", "localhost:5000", "server address to check server status")
var configServer = flag.String("configServiceAddr", "localhost:5001", "server address for Config Service")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*configServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := blodBank.NewBlodBankServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// confID := blodBank.ConfigID{Id: "1"}
	// resp, err := c.GetConfig(ctx, &confID)
	// if err != nil {
	// 	log.Fatalf("Cannot get config: %v", err)
	// }
	//
	// fmt.Println(resp)

	stream, err := c.ListAllConfig(ctx, &blodBank.NoParam{})
	if err != nil {
		log.Fatalf("Cannot get config list: %v", err)
	}

	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListConfigs failed: %v", err)
		}
		fmt.Println("Config: ", config)
	}
}
