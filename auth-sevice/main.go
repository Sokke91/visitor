package main

import (
	"auth/database"
	"auth/models"
	"fmt"
	"net"
	"net/rpc"
)

const rpcPort = "50001"

type Config struct{}

func main() {
	app := Config{}
	database.ConnectToDatabase()
	database.DB.AutoMigrate(&models.Admin{})
	app.rpcListen()
}

func (app *Config) rpcListen() error {
	fmt.Println("Rcp Server gestartet")
	rpcServer := new(RpcServer)
	rpc.Register(rpcServer)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
