package runner

import (
	"context"
	"fmt"
	"os"
	"time"

	cmd "github.com/dracuxan/blod-bank/client/runner/commands"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

func Run(conn *grpc.ClientConn, option string) {
	c := blodBank.NewBlodBankServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch option {
	case "get":
		cmd.GetCommand(ctx, c)

	case "list":
		cmd.ListCommand(ctx, c)

	case "register":
		cmd.RegisterCommand(ctx, c)

	case "update":
		cmd.UpdateCommand(ctx, c)

	case "delete":
		cmd.DeleteCommand(ctx, c)

	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
