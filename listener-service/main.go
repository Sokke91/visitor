package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connectToRabbitMQ()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	consumer, err := NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}
	err = consumer.Listen([]string{"job.LOG", "job.MAIL", "job.ID"})
	if err != nil {
		log.Println(err)
	}
}

func connectToRabbitMQ() (*amqp.Connection, error) {
	var counts int64
	backoff := 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest@localhost")
		if err != nil {
			fmt.Print("rabbitMQ not yet ready")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}
		if counts > 5 {
			return nil, err
		}
		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backoff")
		time.Sleep(backoff)
		continue
	}
	return connection, nil
}
