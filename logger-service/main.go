package main

import (
	"fmt"
	"net"
	"net/rpc"
)

const RpcPort = "50004"

type Config struct{}

func main() {
	app := Config{}
	app.rpcListen()
}

func (app *Config) rpcListen() error {
	fmt.Printf("Starting RPC server on port %s", RpcPort)
	rpcServer := new(RpcServer)
	rpc.Register(rpcServer)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", RpcPort))
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
