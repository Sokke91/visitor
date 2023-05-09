package main

import (
	"auth/database"
	"auth/models"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("hkjsdhdhkurhakjsadjsurhksyrherh")

type Claims struct {
	PersonalNumber string `json:"personaNumber"`
	jwt.StandardClaims
}

func GenerateToken(personalNumber string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		PersonalNumber: personalNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := getToken(tokenString)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return errors.New("invalid Token")
	}
	return nil
}

func CurrentUser(tokenString string) (models.Admin, error) {
	err := ValidateToken(tokenString)
	if err != nil {
		return models.Admin{}, err
	}
	token, _ := getToken(tokenString)
	claims, _ := token.Claims.(*Claims)

	var admin models.Admin
	err = database.DB.Where("personalNumber=?", claims.PersonalNumber).Find(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

func getToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	return token, err
}
