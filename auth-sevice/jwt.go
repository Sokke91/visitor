package main

import (
	"auth/database"
	"auth/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = []byte("hkjsdhdhkurhakjsadjsurhksyrherh")

type Claims struct {
	PersonalNumber string `json:"personaNumber"`
	jwt.StandardClaims
}

func GenerateToken(personalNumber string, ctx *gin.Context) string {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate Token"})
		return ""
	}
	return tokenString
}

func ValidateToken(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return errors.New("invalid Token")
	}
	return nil
}

func CurrentUser(c *gin.Context) (models.Admin, error) {
	err := ValidateToken(c)
	if err != nil {
		return models.Admin{}, err
	}
	token, _ := getToken(c)
	claims, _ := token.Claims.(*Claims)

	var admin models.Admin
	err = database.DB.Where("personalNumber=?", claims.PersonalNumber).Find(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(ctx)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	return token, err
}

func getTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	tokenString := authHeader[7:] // Remove the "Bearer " prefix
	return tokenString
}
