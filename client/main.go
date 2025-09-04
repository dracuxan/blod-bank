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

var addr = flag.String("addr", "localhost:5000", "server address to connect")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := blodBank.NewBlodBankServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// info := &blodBank.Type{
	// 	Name: "Nisarg",
	// 	Type: "B+",
	// }
	// r, err := c.DonateBlod(ctx, info)
	// if err != nil {
	// 	log.Fatalf("unable to donate blod: %v", err)
	// }
	// fmt.Printf("Thxx %s for donating blod!!\n", r.GetName())

	r, err := c.GetBlod(ctx, &blodBank.NoParam{})
	if err != nil {
		log.Fatalf("unable to get blod: %v", err)
	}
	fmt.Printf("Blod samples available\n")
	fmt.Println("\nID\tTYPE")
	fmt.Printf("%s\t%s\n", r.GetId(), r.GetType())
}
