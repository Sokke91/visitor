package controllers

import (
	"broker/database"
	"broker/event"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type CardPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func CreateVisitor(ctx *gin.Context) {
	// TODO CREATE Visitor Logic

	// Create Jobs
	emitter, err := event.NewEventEmmiter(database.RabbitMqConnection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = LogJob(emitter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = SendMailJob(emitter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = CreateCardJob(emitter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func LogJob(e event.Emitter) error {
	p := LogPayload{
		Name: "log",
		Data: "User was Creadet",
	}
	j, err := json.MarshalIndent(&p, "", "\t")
	if err != nil {
		return err
	}
	err = e.Push(string(j), "job.LOG")
	if err != nil {
		return err
	}
	return nil
}

func SendMailJob(e event.Emitter) error {
	p := MailPayload{
		Name: "mail",
		Data: "User was Creadet",
	}
	j, err := json.MarshalIndent(&p, "", "\t")
	if err != nil {
		return err
	}
	err = e.Push(string(j), "job.MAIL")
	if err != nil {
		return err
	}
	return nil
}

func CreateCardJob(e event.Emitter) error {
	p := CardPayload{
		Name: "card",
		Data: "User was creadet",
	}
	j, err := json.MarshalIndent(&p, "", "\t")
	if err != nil {
		return err
	}
	err = e.Push(string(j), "job.ID")
	if err != nil {
		return err
	}
	return nil
}
