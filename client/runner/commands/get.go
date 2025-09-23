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
	"gopkg.in/yaml.v3"
)

func GetCommand(c blodBank.BlodBankServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getID := getCmd.Int64("id", 0, "Config ID to fetch")
	getContentOnly := getCmd.Bool("c", false, "return/display only config content")

	getCmd.Parse(os.Args[2:])

	if *getID == 0 {
		fmt.Println("Need --id=<> flag!!")
		os.Exit(1)
	}
	fmt.Printf("Got config for id %v:\n", *getID)

	conf, err := helper.GetConfig(ctx, &blodBank.ConfigID{Id: *getID}, c)
	if err != nil {
		log.Fatalf("Cannot get config: %v", err)
	}

	if *getContentOnly {
		content, err := yaml.Marshal(conf.Content)
		if err != nil {
			log.Fatalf("cannot marshal response: %v", err)
		}
		fmt.Println(string(content))

	} else {
		orgConf, err := yaml.Marshal(conf)
		if err != nil {
			log.Fatalf("cannot marshal response: %v", err)
		}
		fmt.Println(string(orgConf))
	}
}
