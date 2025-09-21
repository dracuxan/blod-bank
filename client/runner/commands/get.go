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
	getID := getCmd.Int64("id", 0, "Config ID to fetch")

	getCmd.Parse(os.Args[2:])

	if *getID == 0 {
		fmt.Println("Need --id=<> flag!!")
		os.Exit(1)
	}
	fmt.Printf("Got config for id %v:\n", *getID)
	helper.GetConfig(ctx, &blodBank.ConfigID{Id: *getID}, c)
}
