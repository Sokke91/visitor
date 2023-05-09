package controllers

import (
	"broker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAdminDto struct {
	Name           string `json:"name"`
	Prename        string `json:"prename"`
	Password       string `json:"password"`
	PersonalNumber string `json:"personalNumber"`
	Department     string `json:"department"`
}

func CreateAdmin(ctx gin.Context) {
	var input CreateAdminDto
	err := ctx.ShouldBindJSON(&input)
	sendErrorWhenFail(ctx, err)
	admin := models.Admin{
		Name:           input.Name,
		Prename:        input.Prename,
		Password:       input.Password,
		PersonalNumber: input.PersonalNumber,
		Department:     input.Department,
	}
	savedAdmin, err := admin.Save()
	sendErrorWhenFail(ctx, err)
	ctx.JSON(http.StatusOK, gin.H{"data": savedAdmin})
}

func sendErrorWhenFail(ctx gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
