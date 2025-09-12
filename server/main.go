package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
	savedConfigs []*blodBank.ConfigItem
}

var dummyconfig = &blodBank.ConfigItem{
	Id:   "1",
	Name: "msf.conf",
	Content: `
	username: "msf"
	pass: "password"
	`,
	CreatedAt: time.Now().String(),
	UpdatedAt: time.Now().String(),
}

var dummyconfig1 = &blodBank.ConfigItem{
	Id:   "2",
	Name: "shodan.conf",
	Content: `
	username: "shodan"
	pass: "password"
	`,
	CreatedAt: time.Now().String(),
	UpdatedAt: time.Now().String(),
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	blodBank.RegisterBlodBankServiceServer(grpcServer, newServer())

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newServer() *server {
	s := &server{
		savedConfigs: []*blodBank.ConfigItem{dummyconfig, dummyconfig1},
	}
	return s
}

func (s *server) GetConfig(_ context.Context, configItemID *blodBank.ConfigID) (*blodBank.ConfigItem, error) {
	for _, item := range s.savedConfigs {
		if item.Id == configItemID.Id {
			return item, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "invalid id")
}

func (s *server) ListAllConfig(configItem *blodBank.NoParam, stream grpc.ServerStreamingServer[blodBank.ConfigItem]) error {
	for _, item := range s.savedConfigs {
		if err := stream.Send(item); err != nil {
			return err
		}
	}
	return nil
}

// func (s *server) RegisterConfig(ctx context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
// }

// func (s *server) DeleteConfig(ctx context.Context, configItem *blodBank.ConfigID) (*blodBank.Status, error) {
// }
//
// func (s *server) UpdateConfig(ctx context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
// }
