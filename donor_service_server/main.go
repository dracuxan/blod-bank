package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedDonorServiceServer
}

var donors = make(map[string]*blodBank.NewDonor)

func (s *server) RegisterDonor(_ context.Context, newDonor *blodBank.NewDonor) (*blodBank.DonorID, error) {
	id := strconv.Itoa(len(donors) + 1)
	donors[id] = newDonor
	log.Println("New donor registed:", donors[id])
	return &blodBank.DonorID{Id: id}, nil
}

func main() {
	flag.Parse()
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blodBank.RegisterDonorServiceServer(s, &server{})
	log.Printf("server listening on port %v", ls.Addr())
	if err := s.Serve(ls); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
