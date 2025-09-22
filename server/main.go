package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"github.com/dracuxan/blod-bank/server/models"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
	db *gorm.DB
}

func (s *server) RegisterConfig(_ context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	newConf := models.Configs{
		Name:    configItem.Name,
		Content: configItem.Content,
	}

	res := s.db.Create(&newConf)
	if res.Error != nil {
		log.Printf("ERROR inserting config: %v", res.Error)
		return nil, status.Error(codes.Internal, "failed to insert config")
	}

	configItem.Id = newConf.ID

	log.Printf("New config inserted with id: %d", configItem.Id)
	return &blodBank.Status{Status: "Registered new config"}, nil
}

func (s *server) GetConfig(_ context.Context, configItemID *blodBank.ConfigID) (*blodBank.ConfigItem, error) {
	var conf models.Configs

	if err := s.db.First(&conf, configItemID.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "config not found")
	}

	log.Printf("Fetched config with id: %d", configItemID.Id)
	protoConf := &blodBank.ConfigItem{
		Id:        conf.ID,
		Name:      conf.Name,
		Content:   conf.Content,
		CreatedAt: conf.CreatedAt.String(),
		UpdatedAt: conf.UpdatedAt.String(),
	}

	return protoConf, nil
}

func (s *server) ListAllConfig(noParam *blodBank.NoParam, stream grpc.ServerStreamingServer[blodBank.ConfigItem]) error {
	log.Println("streaming list of all the configs")

	var configs []models.Configs
	s.db.Find(&configs)

	for _, i := range configs {
		item := &blodBank.ConfigItem{
			Id:        i.ID,
			Name:      i.Name,
			Content:   i.Content,
			CreatedAt: i.CreatedAt.String(),
			UpdatedAt: i.UpdatedAt.String(),
		}
		if err := stream.Send(item); err != nil {
			return status.Error(codes.Aborted, "bad request")
		}
	}
	return nil
}

func (s *server) DeleteConfig(ctx context.Context, configID *blodBank.ConfigID) (*blodBank.Status, error) {
	var conf models.Configs
	if err := s.db.First(&conf, configID.Id).Error; err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "could not find config")
	}
	if err := s.db.Delete(&conf).Error; err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "failed to delete config")
	}

	log.Printf("Deleted config with id %d", configID.Id)
	return &blodBank.Status{Status: fmt.Sprintf("Deleted config with id: %d", configID.Id)}, nil
}

func (s *server) UpdateConfig(ctx context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	var conf models.Configs
	if err := s.db.First(&conf, configItem.Id).Error; err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "could not find config")
	}
	conf.Name = configItem.Name
	conf.Content = configItem.Content

	if err := s.db.Save(&conf).Error; err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "failed to update config")
	}

	log.Printf("Updated config with id %d", configItem.Id)

	return &blodBank.Status{Status: fmt.Sprintf("updated config with id %d", configItem.Id)}, nil
}

func newServer(db *gorm.DB) *server {
	return &server{db: db}
}

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
	blodBank.RegisterBlodBankServiceServer(grpcServer, newServer(db))

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
