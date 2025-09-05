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

// statusAddr       = flag.String("statusAddr", "localhost:5000", "server address to check server status")
var donorServiceAddr = flag.String("donorServiceAddr", "localhost:5001", "server address for Donor Service")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*donorServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := blodBank.NewDonorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newDonor := blodBank.NewDonor{
		Name:     "Nisarg",
		BlodType: "B+",
	}

	r, err := c.RegisterDonor(ctx, &newDonor)
	if err != nil {
		log.Fatalf("Cannot register donor: %v", err)
	}
	fmt.Printf("Registered donor with id: %s\n", r.GetId())
}
