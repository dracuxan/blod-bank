package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"gopkg.in/yaml.v3"
)

func ListCommand(c blodBank.BlodBankServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
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
