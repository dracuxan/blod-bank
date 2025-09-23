package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"github.com/dracuxan/blod-bank/server/handler"
	"github.com/dracuxan/blod-bank/server/models"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5001, "Server port")

func main() {
	flag.Parse()
	db, err := models.Init()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	blodBank.RegisterBlodBankServiceServer(grpcServer, handler.NewServer(db))

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
