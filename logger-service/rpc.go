package main

import "fmt"

type RpcServer struct{}

type LogPayload struct {
	Entry string `json:"entry"`
}

func (r *RpcServer) LogEntry(payload LogPayload, resp *string) error {
	fmt.Println("Receive Log RPC and start logging")
	*resp = "Log entry created"
	return nil
}
