package jwtclient

import (
	"github.com/golang-jwt/jwt"
)

type JWTClient interface {
	generate(userID string) (string, error)
}

type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

// Generating token with CustomClaims payload
func (cc *CustomClaims) generate(JWTSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	ss, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func GenerateToken(jwtC JWTClient, JWTSecret string) (string, error) {
	token, err := jwtC.generate(JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
