package main

import "github.com/gin-gonic/gin"

func getTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	tokenString := authHeader[7:] // Remove the "Bearer " prefix
	return tokenString
}
