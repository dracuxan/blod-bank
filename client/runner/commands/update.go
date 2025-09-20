package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func UpdateCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	updCmd := flag.NewFlagSet("register", flag.ExitOnError)
	updID := updCmd.String("id", "", "id of the config")
	updName := updCmd.String("name", "", "name of config")
	updContent := updCmd.String("content", "", "content")

	updCmd.Parse(os.Args[2:])

	if *updName == "" || *updContent == "" {
		fmt.Println("Need --id=<> --name=<> and --content=<> flag!!")
		os.Exit(1)
	}

	fmt.Println("Adding new config...")

	newConf := blodBank.ConfigItem{
		Id:      *updID,
		Name:    *updName,
		Content: *updContent,
	}
	helper.UpdateConfig(ctx, &newConf, c)
}
