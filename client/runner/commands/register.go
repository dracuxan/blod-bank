package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func RegisterCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	regName := registerCmd.String("name", "", "name of config")
	regContent := registerCmd.String("content", "", "content")
	regFile := registerCmd.String("file", "", "location of the config file")

	registerCmd.Parse(os.Args[2:])

	if *regName == "" {
		fmt.Println("Need --name=<> flag!!")
		os.Exit(1)
	}

	fmt.Println("Adding new config...")

	if *regContent != "" {
		newConf := blodBank.ConfigItem{
			Name:    *regName,
			Content: *regContent,
		}
		helper.RegisterConfig(ctx, &newConf, c)
	} else if *regFile != "" {
		content, err := os.ReadFile(*regFile)
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		newConf := blodBank.ConfigItem{
			Name:    *regName,
			Content: string(content),
		}
		helper.RegisterConfig(ctx, &newConf, c)
	} else {
		fmt.Println("Need --content or --file!!")
		os.Exit(1)
	}
}
