package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
type LogPayload struct {
	Entry string `json:"entry"`
}

type MailPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type CardPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Println(d.RoutingKey)
			switch d.RoutingKey {
			case "job.MAIL":
				var payload MailPayload
				_ = json.Unmarshal(d.Body, &payload)
				go handleMailJob(payload)
			case "job.LOG":
				var payload LogPayload
				_ = json.Unmarshal(d.Body, &payload)
				go handleLoggingJob(payload)
			case "job.ID":
				var payload CardPayload
				_ = json.Unmarshal(d.Body, &payload)
				go handleCreateVisitorIdJob(payload)
			}
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	<-forever

	return nil
}

func handleMailJob(payload MailPayload) {
	fmt.Println("Call Mail Service")
}

func handleCreateVisitorIdJob(payload CardPayload) {
	fmt.Println("Call Creade VisitorId Service")
}

func handleLoggingJob(payload LogPayload) {
	fmt.Println("Call Log Service")
	client, err := rpc.Dial("tcp", "localhost:50004")
	if err != nil {
		fmt.Println(err)
	}
	var response *string
	err = client.Call("RpcServer.LogEntry", payload, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func logEvent(entry Payload) error {
	fmt.Println("Log event")
	return nil
}
