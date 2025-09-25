package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"gopkg.in/yaml.v3"
)

func GetCommand(c blodBank.BlodBankServiceClient) {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getID := getCmd.Int64("id", 0, "Config ID to fetch")
	getContentOnly := getCmd.Bool("c", false, "return/display only config content")

	getCmd.Parse(os.Args[2:])

	if *getID == 0 {
		helper.CheckMissingFlag("--id")
	}

	fmt.Printf("Getting config for id %v:\n", *getID)

	conf, err := helper.GetConfig(&blodBank.ConfigID{Id: *getID}, c)
	helper.CheckCommonError(err, "Cannot get config")

	if *getContentOnly {
		content, err := yaml.Marshal(conf.Content)
		helper.CheckMarshalError(err)
		fmt.Println(string(content))
		os.Exit(0)
	}

	orgConf, err := yaml.Marshal(conf)
	helper.CheckMarshalError(err)
	fmt.Println(string(orgConf))
}
