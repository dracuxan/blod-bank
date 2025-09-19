package runner

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

func Run(conn *grpc.ClientConn, option string) {
	// options for get command
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getID := getCmd.String("id", "", "Config ID to fetch")

	// options for register command
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	regName := registerCmd.String("name", "", "name of config")
	regContent := registerCmd.String("content", "", "content")
	regFile := registerCmd.String("file", "", "location of the config file")

	// options for delete command
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.String("id", "", "id of config to delete")

	// options for update command
	updCmd := flag.NewFlagSet("register", flag.ExitOnError)
	updID := updCmd.String("id", "", "id of the config")
	updName := updCmd.String("name", "", "name of config")
	updContent := updCmd.String("content", "", "content")

	c := blodBank.NewBlodBankServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch option {
	case "get":
		getCmd.Parse(os.Args[2:])
		if *getID == "" {
			fmt.Println("Need --id=<> flag!!")
			os.Exit(1)
		}
		fmt.Printf("Got config for id %v:\n", *getID)
		helper.GetConfig(ctx, &blodBank.ConfigID{Id: *getID}, c)

	case "list":
		fmt.Println("Listing all configs:")
		helper.ListAllConfig(ctx, c)

	case "register":
		registerCmd.Parse(os.Args[2:])
		if *regName == "" {
			fmt.Println("Need --name=<> flag!!")
			os.Exit(1)
		}

		fmt.Println("Adding new config...")

		if *regContent != "" {
			newConf := blodBank.ConfigItem{
				Name:    *regName,
				Content: *regContent,
			}
			helper.RegisterConfig(ctx, &newConf, c)
		} else if *regFile != "" {
			content, err := os.ReadFile(*regFile)
			if err != nil {
				log.Fatalf("failed to read file: %v", err)
			}

			newConf := blodBank.ConfigItem{
				Name:    *regName,
				Content: string(content),
			}
			helper.RegisterConfig(ctx, &newConf, c)
		} else {
			fmt.Println("Need --content or --file!!")
			os.Exit(1)
		}

	case "update":
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

	case "delete":
		delCmd.Parse(os.Args[2:])
		if *delID == "" {
			fmt.Println("neeeed --id flag!!")
			os.Exit(1)
		}
		fmt.Println("deleting config with id:", *delID)
		helper.DeleteConfig(ctx, &blodBank.ConfigID{Id: *delID}, c)

	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
