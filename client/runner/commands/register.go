package cmd

import (
	"flag"
	"fmt"
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
		helper.CheckMissingFlag("--name")
	}

	fmt.Println("Adding new config...")

	if *regContent != "" {
		newConf := &blodBank.ConfigItem{
			Name:    *regName,
			Content: *regContent,
		}

		status, err := helper.RegisterConfig(newConf, c)
		helper.CheckCommonError(err, "Cannot register config")

		fmt.Println(status)
	} else if *regFile != "" {
		newConf, err := helper.CreateConfig(*regName, *regFile)
		helper.CheckCommonError(err, "failed to read file")

		status, err := helper.RegisterConfig(newConf, c)
		helper.CheckCommonError(err, "Cannot register config")

		fmt.Println(status)
	} else if *editor {
		// use editor to create file (vim)
		filename, err := helper.OpenEditor(*regName)
		helper.CheckCommonError(err, "Unable to edit/create file:")

		newConf, err := helper.CreateConfig(*regName, filename)
		helper.CheckCommonError(err, "failed to read file")

		status, err := helper.RegisterConfig(newConf, c)
		helper.CheckCommonError(err, "Cannot register config")

		fmt.Println(status)
	} else {
		helper.CheckMissingFlag("--content or --file")
	}
}
