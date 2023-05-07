package main

import (
	"broker/controllers"
	"broker/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	connectToDabase()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env konnte nicht geladen werden")
	}
}

func connectToDabase() {
	database.ConnectToDatabase()
}

func startServer() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.Run(":3000")
}
