package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dracuxan/blod-bank/client/helper"
	blodBank "github.com/dracuxan/blod-bank/proto"
	"github.com/dracuxan/blod-bank/server/handler"
	"github.com/dracuxan/blod-bank/server/models"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5000, "Server port")

func main() {
	flag.Parse()
	db, err := models.Init()
	helper.CheckCommonError(err, "failed to connect to db")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	helper.CheckCommonError(err, "failed to listen")

	grpcServer := grpc.NewServer()
	blodBank.RegisterBlodBankServiceServer(grpcServer, handler.NewServer(db))

	log.Printf("gRPC server listening on %v", lis.Addr())
	err = grpcServer.Serve(lis)
	helper.CheckCommonError(err, "failed to serve")
}
