package helper

import (
	"context"
	"io"
	"log"

	blodBank "github.com/dracuxan/blod-bank/proto"
)

func GetConfig(ctx context.Context, id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) (*blodBank.ConfigItem, error) {
	resp, err := c.GetConfig(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ListAllConfig(ctx context.Context, c blodBank.BlodBankServiceClient) ([]*blodBank.ConfigItem, error) {
	stream, err := c.ListAllConfig(ctx, &blodBank.NoParam{})
	var configs []*blodBank.ConfigItem

	if err != nil {
		log.Fatalf("Cannot get config list: %v", err)
	}

	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListConfigs failed: %v", err)
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func RegisterConfig(ctx context.Context, item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) (string, error) {
	status, err := c.RegisterConfig(ctx, item)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}

func DeleteConfig(ctx context.Context, id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) (string, error) {
	status, err := c.DeleteConfig(ctx, id)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}

func UpdateConfig(ctx context.Context, item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) (string, error) {
	status, err := c.UpdateConfig(ctx, item)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}
