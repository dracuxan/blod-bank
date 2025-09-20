package cmd

import (
	"context"
	"fmt"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

func ListCommand(ctx context.Context, c blodBank.BlodBankServiceClient) {
	fmt.Println("Listing all configs:")
	helper.ListAllConfig(ctx, c)
}
