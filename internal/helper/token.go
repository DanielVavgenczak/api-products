package helper

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id string, expired int) (string, error)  {
	tokenConfig :=  os.Getenv("TOKEN_CONFIG")
	expirationTime := time.Now().Add(time.Second * time.Duration(expired)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id 
	claims["exp"] = expirationTime
	
	tokenGenrated, err := token.SignedString([]byte(tokenConfig))
	if err != nil {
		return "", err
	}

	return tokenGenrated, nil
}


func ValidToken(tokenBearer string) error {
	tokenConfig :=  os.Getenv("TOKEN_CONFIG")
	fmt.Println("valida", tokenConfig)
	token, err := jwt.Parse(tokenBearer,func (t *jwt.Token) (interface{}, error)  {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok  {
				return nil, errors.New("token not authorized")
			}
			return []byte(tokenConfig), nil
	})
	if err != nil {
		fmt.Println("Aqui sera")
		return errors.New("token is invalid")
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("token is invalid")
	}

	return nil
}