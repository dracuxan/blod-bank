package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
	savedConfigs map[int]*blodBank.ConfigItem
}

func (s *server) GetConfig(_ context.Context, configItemID *blodBank.ConfigID) (*blodBank.ConfigItem, error) {
	for _, item := range s.savedConfigs {
		if item.Id == configItemID.Id {
			log.Printf("sending config with id %s", configItemID.Id)
			return item, nil
		}
	}

	log.Printf("ERROR! cannot find config with id %s", configItemID.Id)
	return nil, status.Errorf(codes.NotFound, "invalid id")
}

func (s *server) ListAllConfig(configItem *blodBank.NoParam, stream grpc.ServerStreamingServer[blodBank.ConfigItem]) error {
	log.Println("streaming list of all the configs")

	for _, item := range s.savedConfigs {
		if err := stream.Send(item); err != nil {
			return status.Error(codes.Aborted, "bad request")
		}
	}

	return nil
}

func (s *server) RegisterConfig(_ context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	id := len(s.savedConfigs) + 1

	newConfig := &blodBank.ConfigItem{
		Id:        strconv.Itoa(id),
		Name:      configItem.Name,
		Content:   configItem.Content,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	s.savedConfigs[id] = newConfig

	log.Printf("New config with id %d registered", id)

	return &blodBank.Status{Status: "Registerd new config"}, nil
}

func (s *server) DeleteConfig(ctx context.Context, configID *blodBank.ConfigID) (*blodBank.Status, error) {
	id, err := strconv.Atoi(configID.Id)
	if err != nil {
		log.Printf("ERROR while deleteing cofig with id %s: invalid id", configID)
		return nil, status.Error(codes.Aborted, "bad request. invalid id")
	}
	delete(s.savedConfigs, id)

	log.Printf("Deleted config with id %d", id)
	return &blodBank.Status{Status: fmt.Sprintf("Deleted config with id: %d", id)}, nil
}

func (s *server) UpdateConfig(ctx context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	id, err := strconv.Atoi(configItem.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, "bad request. invalid id")
	}
	_, ok := s.savedConfigs[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "invalid id")
	}

	s.savedConfigs[id] = configItem
	log.Printf("Updated config with id %d", id)

	return &blodBank.Status{Status: fmt.Sprintf("updated config with id %d", id)}, nil
}

func newServer() *server {
	s := &server{
		savedConfigs: map[int]*blodBank.ConfigItem{},
	}
	return s
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
