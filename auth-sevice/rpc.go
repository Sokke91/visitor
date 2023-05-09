package main

import (
	"auth/database"
	"auth/models"
	"errors"
)

type RpcServer struct{}

type Credentials struct {
	PersonalNumber string
	Password       string
}

func (r *RpcServer) Login(credentials Credentials, result *string) error {
	var admin models.Admin
	err := database.DB.Where("personalNumber=?", credentials.PersonalNumber).Find(&admin).Error
	if err != nil {
		return err
	}
	if admin.Password != credentials.Password {
		return errors.New("invalid password")
	}
	*result, err = GenerateToken(credentials.PersonalNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *RpcServer) CheckToken(token string) error {
	err := ValidateToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (r *RpcServer) CurrentUser(token string) error {
	return nil
}
