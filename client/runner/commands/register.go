package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func RegisterCommand(c blodBank.BlodBankServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
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
		status, err := helper.RegisterConfig(ctx, &newConf, c)
		if err != nil {
			log.Fatalf("Cannot register config: %v", err)
		}
		fmt.Println(status)
	} else if *regFile != "" {
		content, err := os.ReadFile(*regFile)
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}
		newConf := blodBank.ConfigItem{
			Name:    *regName,
			Content: string(content),
		}
		status, err := helper.RegisterConfig(ctx, &newConf, c)
		if err != nil {
			log.Fatalf("Cannot register config: %v", err)
		}
		fmt.Println(status)
	} else {
		fmt.Println("Need --content or --file!!")
		os.Exit(1)
	}
}
