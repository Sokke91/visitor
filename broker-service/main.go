package main

import (
	"broker/controllers"
	"broker/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	connectToDabase()
	startServer()
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
	fmt.Println("Starte Broker on Port 3000")
	r := gin.Default()
	r.POST("/login", controllers.Login)
	r.Run(":3000")
}
