package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	blodBank "github.com/dracuxan/blod-bank/proto"
)

func GetConfig(ctx context.Context, id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) {
	resp, err := c.GetConfig(ctx, id)
	if err != nil {
		log.Fatalf("Cannot get config: %v", err)
	}
	orgConf, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Fatalf("cannot marshal response: %v", err)
	}
	fmt.Println(string(orgConf))
}

func ListAllConfig(ctx context.Context, c blodBank.BlodBankServiceClient) {
	stream, err := c.ListAllConfig(ctx, &blodBank.NoParam{})
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
		}
		orgConf, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatalf("cannot marshal response: %v", err)
		}
		fmt.Println(string(orgConf))
	}
}

func RegisterConfig(ctx context.Context, item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) {
	status, err := c.RegisterConfig(ctx, item)
	if err != nil {
		log.Fatalf("Cannot register config: %v", err)
	}
	fmt.Println(status.Status)
}

func DeleteConfig(ctx context.Context, id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) {
	status, err := c.DeleteConfig(ctx, id)
	if err != nil {
		log.Fatalf("Cannot delete config: %v", err)
	}
	fmt.Println(status.Status)
}

func UpdateConfig(ctx context.Context, item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) {
	status, err := c.UpdateConfig(ctx, item)
	if err != nil {
		log.Fatalf("Cannot update config: %v", err)
	}
	fmt.Println(status.Status)
}
