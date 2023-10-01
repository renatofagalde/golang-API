package model

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang-basic/config/rest_err"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToen() (string, *rest_err.RestErr) {

	secretKey := os.Getenv(JWT_SECRET_KEY)

	//quais campos dentro do jwt
	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("erro ao gerar jwt token, err=%s", err.Error()))
	}
	return tokenString, nil
}
