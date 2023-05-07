package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOk, gin.H{"data": "Pong"})
}
