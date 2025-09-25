package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func DeleteCommand(c blodBank.BlodBankServiceClient) {
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.Int64("id", 0, "id of config to delete")
	delCmd.Parse(os.Args[2:])

	if *delID == 0 {
		helper.CheckMissingFlag("--id")
	}

	fmt.Println("deleting config with id:", *delID)

	status, err := helper.DeleteConfig(&blodBank.ConfigID{Id: *delID}, c)
	helper.CheckCommonError(err, "cannot delete config")

	fmt.Println(status)
}
