package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret []byte, id int64, exp time.Duration) (string, error) {
	cl := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(exp * time.Second).Unix(),
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
	sgTok, err := tok.SignedString(secret)
	if err != nil {
		log.Println("GenerateToken: ", err.Error())
		return "", err
	}

	return sgTok, nil
}

func ValidateToken(tokStr string, secret []byte) (string, error) {
	tok, err := jwt.Parse(tokStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected error")
		}

		return secret, nil
	})

	if err != nil {
		log.Println("ValidateToken: ", err.Error())
		return "", err
	}

	if cl, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		return fmt.Sprintf("%v", cl["user_id"]), nil
	}

	return "", errors.New("invalid token")
}
