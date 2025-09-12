package main

import (
	"context"
	"flag"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 5001, "Server port")

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
	savedConfigs []*blodBank.ConfigItem
}

func (s *server) GetConfig(_ context.Context, configItemID *blodBank.ConfigID) (*blodBank.ConfigItem, error) {
	for _, item := range s.savedConfigs {
		if item.Id == configItemID.Id {
			return item, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "invalid id")
}

func (s *server) ListAllConfig(ctx context.Context, configItem *blodBank.NoParam, stream blodBank.BlodBankService_ListAllConfigServer) error {
	for _, item := range s.savedConfigs {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := stream.Send(item); err != nil {
				return err
			}
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
