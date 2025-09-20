package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func GetCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getID := getCmd.String("id", "", "Config ID to fetch")

	getCmd.Parse(os.Args[2:])

	if *getID == "" {
		fmt.Println("Need --id=<> flag!!")
		os.Exit(1)
	}
	fmt.Printf("Got config for id %v:\n", *getID)
	helper.GetConfig(ctx, &blodBank.ConfigID{Id: *getID}, c)
}
