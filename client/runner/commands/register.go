package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func RegisterCommand(c blodBank.BlodBankServiceClient) {
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	regName := registerCmd.String("name", "", "name of config")
	regContent := registerCmd.String("content", "", "content")
	regFile := registerCmd.String("file", "", "location of the config file")
	editor := registerCmd.Bool("editor", false, "Use editor to create a config file")

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

		status, err := helper.RegisterConfig(&newConf, c)
		if err != nil {
			log.Fatalf("Cannot register config: %v", err)
		}

		fmt.Println(status)
	} else if *regFile != "" {
		newConf, err := helper.CreateConfig(*regName, *regFile)
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		status, err := helper.RegisterConfig(newConf, c)
		if err != nil {
			log.Fatalf("Cannot register config: %v", err)
		}

		fmt.Println(status)
	} else if *editor {
		// use editor to create file (vim)
		filename, err := helper.OpenEditor(*regName)
		if err != nil {
			log.Fatalf("Unable to edit/create file: %v", err)
		}

		newConf, err := helper.CreateConfig(*regName, filename)
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		status, err := helper.RegisterConfig(newConf, c)
		if err != nil {
			log.Fatalf("Cannot register config: %v", err)
		}

		fmt.Println(status)
	} else {
		fmt.Println("Need --content or --file!!")
		os.Exit(1)
	}
}
