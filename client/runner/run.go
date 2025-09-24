package runner

import (
	"fmt"
	"os"

	cmd "github.com/dracuxan/blod-bank/client/runner/commands"
	"github.com/dracuxan/blod-bank/client/runner/rpc"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"google.golang.org/grpc"
)

func Run(conn *grpc.ClientConn, option string) {
	c := blodBank.NewBlodBankServiceClient(conn)

	switch option {
	case "get":
		cmd.GetCommand(c)

	case "list":
		cmd.ListCommand(c)

	case "register":
		cmd.RegisterCommand(c)

	case "update":
		cmd.UpdateCommand(c)

	case "delete":
		cmd.DeleteCommand(c)
	case "server":
		rpc.RpcServer(c)

	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
