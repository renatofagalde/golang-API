package model

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang-basic/config/rest_err"
	"os"
	"strings"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

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

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secretKey := os.Getenv(JWT_SECRET_KEY)
	token, err := jwt.Parse(RemoveBearer(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}
		return nil, rest_err.NewBadRequestError("token inválido")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("token inválido")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
	}, nil
}

func RemoveBearer(token string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(token, prefix) {
		token = strings.TrimPrefix(prefix, token)
	}
	return token
}
