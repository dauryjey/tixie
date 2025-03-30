package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	UserID string
	Email  string
}

func GenerateJWT(payload JWTPayload) (*string, error) {
	appEnv := GetEnv("APP_ENV")

	var (
		key         []byte
		token       *jwt.Token
		signedToken string
		issuer      string
	)

	if appEnv == "production" {
		issuer = ""
	} else {
		issuer = fmt.Sprintf("http://localhost:%s", GetEnv("AUTH_PORT"))
	}

	key = []byte(GetEnv("JWT_SECRET"))

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   issuer,
		"sub":   payload.UserID,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Unix(),
		"email": payload.Email,
	})

	signedToken, err := token.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &signedToken, err
}
