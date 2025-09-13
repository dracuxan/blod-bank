package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/dracuxan/blod-bank/client/helper"
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
	// helper.DeleteConfig(ctx, &confID, c)
	// helper.GetConfig(ctx, &confID, c)
	// println()

	helper.ListAllConfig(ctx, c)

	// dummyconfig := &blodBank.ConfigItem{
	// 	Name: "msf.conf",
	// 	Content: `username: "msf"
	//  pass: "password"
	//  	`,
	// }
	//
	// helper.RegisterConfig(ctx, dummyconfig, c)
}
