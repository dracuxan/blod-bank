package cmd

import (
	"fmt"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"gopkg.in/yaml.v3"
)

func ListCommand(c blodBank.BlodBankServiceClient) {
	fmt.Println("Listing all configs:")

	allConfigs, err := helper.ListAllConfig(c)
	helper.CheckCommonError(err, "cannot list config")

	for _, config := range allConfigs {
		orgConf, err := yaml.Marshal(config)
		helper.CheckMarshalError(err)
		fmt.Println(string(orgConf))
	}
}
