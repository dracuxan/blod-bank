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

func DeleteCommand(c blodBank.BlodBankServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.Int64("id", 0, "id of config to delete")
	delCmd.Parse(os.Args[2:])
	if *delID == 0 {
		fmt.Println("neeeed --id flag!!")
		os.Exit(1)
	}
	fmt.Println("deleting config with id:", *delID)
	helper.DeleteConfig(ctx, &blodBank.ConfigID{Id: *delID}, c)
	status, err := helper.DeleteConfig(ctx, &blodBank.ConfigID{Id: *delID}, c)
	if err != nil {
		log.Fatalf("Cannot delete config: %v", err)
	}
	fmt.Println(status)
}
