package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func UpdateCommand(c blodBank.BlodBankServiceClient) {
	updCmd := flag.NewFlagSet("register", flag.ExitOnError)
	updID := updCmd.Int64("id", 0, "id of the config")

	updName := updCmd.String("name", "", "name of config")
	updContent := updCmd.String("content", "", "content")
	updFile := updCmd.String("file", "", "location of the config file")

	updCmd.Parse(os.Args[2:])

	if *updName == "" {
		helper.CheckMissingFlag("--id")
	}

	fmt.Println("Adding new config...")

	if *updContent != "" {
		newConf := &blodBank.ConfigItem{
			Id:      *updID,
			Name:    *updName,
			Content: *updContent,
		}
		status, err := helper.UpdateConfig(newConf, c)
		helper.CheckCommonError(err, "Cannot update config")

		fmt.Println(status)
	} else if *updFile != "" {
		content, err := os.ReadFile(*updFile)
		helper.CheckCommonError(err, "failed to read file")

		newConf := &blodBank.ConfigItem{
			Id:      *updID,
			Name:    *updName,
			Content: string(content),
		}

		status, err := helper.UpdateConfig(newConf, c)
		helper.CheckCommonError(err, "Cannot update config")

		fmt.Println(status)
	} else {
		helper.CheckMissingFlag("--content or --file")
	}
}
