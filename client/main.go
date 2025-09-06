package main

import (
	"context"
	"encoding/json"
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

	donorID := blodBank.DonorID{Id: "2"}
	nr, err := c.GetDonor(ctx, &donorID)
	if err != nil {
		log.Fatalf("Cannot get donor: %v", err)
	}

	fmt.Printf("Got Donor Info: %v\n", nr)

	resp, err := c.GetAllDonors(ctx, &blodBank.NoParam{})
	if err != nil {
		log.Fatalf("Cannot get donor list: %v", err)
	}
	orgList, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(orgList))
}
