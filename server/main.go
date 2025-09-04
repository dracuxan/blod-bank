package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5000, "Server port")

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
}

func (s *server) GetBlod(_ context.Context, in *blodBank.NoParam) (*blodBank.Samples, error) {
	log.Printf("Sending Blod Samples.")
	return &blodBank.Samples{
		Id:   "1",
		Type: "B+",
	}, nil
}

func main() {
	flag.Parse()
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blodBank.RegisterBlodBankServiceServer(s, &server{})
	log.Printf("server listening on port %v", ls.Addr())
	if err := s.Serve(ls); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
