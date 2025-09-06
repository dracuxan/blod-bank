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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedDonorServiceServer
}

var donors = make(map[int]*blodBank.DonorInfo)

func (s *server) GetAllDonors(_ context.Context, in *blodBank.NoParam) (*blodBank.DonorList, error) {
	var allDonors []*blodBank.DonorInfo
	for _, i := range donors {
		allDonors = append(allDonors, i)
	}
	log.Println("Sending donor list")
	return &blodBank.DonorList{Donors: allDonors}, nil
}

func (s *server) GetDonor(_ context.Context, in *blodBank.DonorID) (*blodBank.DonorInfo, error) {
	donorID, _ := strconv.Atoi(in.GetId())
	donor, ok := donors[donorID]

	if ok != true {
		return nil, status.Errorf(codes.NotFound, "donor with id %d not found", donorID)
	}
	log.Printf("Sending info for donor %s", donor.Name)
	return &blodBank.DonorInfo{Id: in.GetId(), Name: donor.Name, BlodType: donor.BlodType}, nil
}

func (s *server) RegisterDonor(_ context.Context, donor *blodBank.NewDonor) (*blodBank.DonorID, error) {
	id := len(donors) + 1
	newDonor := &blodBank.DonorInfo{
		Id:       strconv.Itoa(id),
		Name:     donor.Name,
		BlodType: donor.BlodType,
	}
	donors[id] = newDonor
	log.Println("New donor registed:", donors[id])
	return &blodBank.DonorID{Id: strconv.Itoa(id)}, nil
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
