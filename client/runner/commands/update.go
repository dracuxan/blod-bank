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

func UpdateCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	updCmd := flag.NewFlagSet("register", flag.ExitOnError)
	updID := updCmd.Int64("id", 0, "id of the config")

	updName := updCmd.String("name", "", "name of config")
	updContent := updCmd.String("content", "", "content")
	updFile := updCmd.String("file", "", "location of the config file")

	updCmd.Parse(os.Args[2:])

	if *updName == "" {
		fmt.Println("Need --id=<> --name=<> and (--content=<> or --file) flag!!")
		os.Exit(1)
	}

	fmt.Println("Adding new config...")

	if *updContent != "" {
		newConf := blodBank.ConfigItem{
			Id:      *updID,
			Name:    *updName,
			Content: *updContent,
		}
		status, err := helper.UpdateConfig(ctx, &newConf, c)
		if err != nil {
			log.Fatalf("Cannot update config: %v", err)
		}
		fmt.Println(status)
	} else if *updFile != "" {
		content, err := os.ReadFile(*updFile)
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}
		newConf := blodBank.ConfigItem{
			Id:      *updID,
			Name:    *updName,
			Content: string(content),
		}
		status, err := helper.UpdateConfig(ctx, &newConf, c)
		if err != nil {
			log.Fatalf("Cannot update config: %v", err)
		}
		fmt.Println(status)
	} else {
		fmt.Println("Need --content or --file!!")
		os.Exit(1)
	}
}
