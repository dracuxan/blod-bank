package main

import (
	"fmt"
	"log"
	"net/rpc"

	"gopkg.in/yaml.v3"
)

type config struct {
	Id        int64
	Name      string
	Content   string
	CreatedAt string
	UpdatedAt string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer client.Close()

	var resp []config
	err = client.Call("Server.ListAll", struct{}{}, &resp)
	if err != nil {
		log.Fatal("error calling ListAll:", err)
	}

	for _, config := range resp {
		orgConf, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("cannot marshal response: %v", err)
		}
		fmt.Println(string(orgConf))
	}
}
