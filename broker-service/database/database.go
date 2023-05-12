package database

import (
	"fmt"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RabbitMqConnection *amqp.Connection

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	DB = db
	fmt.Println("Connect to dabase")
}

func ConnectToRabbitMQ() {
	rabbitConn, err := connecWithBackoff()
	if err != nil {
		fmt.Println("RabbitMQ is not ready. Exit Service")
		os.Exit(1)
	}
	RabbitMqConnection = rabbitConn
}

func connecWithBackoff() (*amqp.Connection, error) {
	var count int64
	backOff := 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("Could not connect to rabbitmq")
			count++
		} else {
			fmt.Println("Connected to rabbitMQ")
			connection = c
			break
		}

		if count > 5 {
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		fmt.Println("Backoff...")
		time.Sleep(backOff)
		continue

	}
	return connection, nil
}
