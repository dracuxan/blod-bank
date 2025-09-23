package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"gopkg.in/yaml.v3"
)

func ListCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	fmt.Println("Listing all configs:")

	allConfigs, err := helper.ListAllConfig(ctx, c)
	if err != nil {
		log.Fatalf("client.ListConfigs failed: %v", err)
	}
	for _, config := range allConfigs {
		orgConf, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("cannot marshal response: %v", err)
		}
		fmt.Println(string(orgConf))
	}
}
