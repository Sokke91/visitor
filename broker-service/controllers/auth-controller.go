package controllers

import (
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	PersonalNumber string
	Password       string
}

func Login(ctx *gin.Context) {
	client, err := rpc.Dial("tcp", "localhost:50001")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var input Credentials
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var response string
	err = client.Call("RpcServer.Login", input, &response)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
