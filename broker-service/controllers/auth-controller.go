package controllers

import (
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string
	Password string
}

func Login(ctx *gin.Context) {
	client, err := rpc.Dial("tcp", "localhost:50001")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var response string
	credentials := Credentials{
		Username: "di32655",
		Password: "geheim",
	}
	err = client.Call("RpcServer.Login", credentials, &response)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
