package runner

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

func Run(conn *grpc.ClientConn) {
	c := blodBank.NewBlodBankServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch os.Args[1] {
	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		id := getCmd.String("id", "", "Config ID to fetch")
		if len(os.Args[2:]) != 0 {
			getCmd.Parse(os.Args[2:])
		} else {
			fmt.Println("Need --id=<> flag!!")
			os.Exit(1)
		}
		fmt.Printf("Got config for id %v:\n", *id)
		helper.GetConfig(ctx, &blodBank.ConfigID{Id: *id}, c)

	case "list":
		fmt.Println("Listing all configs:")
		helper.ListAllConfig(ctx, c)

	case "register":
		registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
		name := registerCmd.String("name", "", "name of config")
		content := registerCmd.String("content", "", "content")
		if len(os.Args[2:]) != 0 {
			registerCmd.Parse(os.Args[2:])
		} else {
			fmt.Println("Need --name=<> and --content=<> flag!!")
			os.Exit(1)
		}
		fmt.Println("Adding new config...")
		newConf := blodBank.ConfigItem{
			Name:    *name,
			Content: *content,
		}
		helper.RegisterConfig(ctx, &newConf, c)

	case "update":
		updCmd := flag.NewFlagSet("register", flag.ExitOnError)

		id := updCmd.String("id", "", "id of the config")
		name := updCmd.String("name", "", "name of config")
		content := updCmd.String("content", "", "content")

		if len(os.Args[2:]) != 0 {
			updCmd.Parse(os.Args[2:])
		} else {
			fmt.Println("Need --id=<> --name=<> and --content=<> flag!!")
			os.Exit(1)
		}

		fmt.Println("Adding new config...")

		newConf := blodBank.ConfigItem{
			Id:      *id,
			Name:    *name,
			Content: *content,
		}
		helper.UpdateConfig(ctx, &newConf, c)

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.String("id", "", "id of config to delete")
		if len(os.Args[2:]) != 0 {
			delCmd.Parse(os.Args[2:])
		} else {
			fmt.Println("neeeed --id flag!!")
			os.Exit(1)
		}

		fmt.Println("deleting config with id:")
		helper.DeleteConfig(ctx, &blodBank.ConfigID{Id: *id}, c)

	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
