package main

import (
	"errors"
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := getTokenFromRequest(ctx)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("not authorized")})
			ctx.Abort()
			return
		}
		err := callCheckTokenService(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("not authorized")})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func callCheckTokenService(tokenString string) error {
	var resp string // no response needed...
	client, err := rpc.Dial("tcp", "localhost:50001")
	if err != nil {
		return err
	}
	err = client.Call("RpcServer.CheckToken", tokenString, &resp)
	if err != nil {
		return err
	}
	return nil
}
