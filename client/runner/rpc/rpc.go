package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
)

type Server struct {
	c blodBank.BlodBankServiceClient
}

type Args struct {
	ID int64
}

type Status struct {
	status string
}

type Config struct {
	ID        int64
	Name      string
	Content   string
	CreatedAt string
	UpdatedAt string
}

func (s *Server) ListAll(_ struct{}, resp *[]Config) error {
	allConfigs, err := helper.ListAllConfig(s.c)
	if err != nil {
		return err
	}
	var results []Config
	for _, conf := range allConfigs {
		results = append(results, Config{
			ID:        conf.Id,
			Name:      conf.Name,
			Content:   conf.Content,
			CreatedAt: conf.CreatedAt,
			UpdatedAt: conf.UpdatedAt,
		})
	}
	*resp = results
	return nil
}

func (s *Server) GetConfig(args *Args, resp *Config) error {
	conf, err := helper.GetConfig(&blodBank.ConfigID{Id: args.ID}, s.c)
	if err != nil {
		return err
	}
	var getConf Config
	getConf = Config{
		ID:        conf.Id,
		Name:      conf.Name,
		Content:   conf.Content,
		CreatedAt: conf.CreatedAt,
		UpdatedAt: conf.UpdatedAt,
	}
	*resp = getConf
	return nil
}

func RpcServer(c blodBank.BlodBankServiceClient) {
	newServer := &Server{c: c}
	rpc.Register(newServer)

	lis, err := net.Listen("tcp", ":9090")
	helper.CheckCommonError(err, "Error listening")
	defer lis.Close()

	fmt.Println("Server is listening on port 9090...")
	rpc.Accept(lis)
}
