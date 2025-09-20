package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func DeleteCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.String("id", "", "id of config to delete")
	delCmd.Parse(os.Args[2:])
	if *delID == "" {
		fmt.Println("neeeed --id flag!!")
		os.Exit(1)
	}
	fmt.Println("deleting config with id:", *delID)
	helper.DeleteConfig(ctx, &blodBank.ConfigID{Id: *delID}, c)
}

