package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	jwtclient "github.com/mochganjarn/go-template-project/external/jwt_client"
)

func JwtTokenGenerator(secret *ClientConnection, userID uint) (string, error) {
	CurrenTime := time.Now()
	ExpTime := CurrenTime.Add(24 * time.Hour)
	claims := jwtclient.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpTime.Unix(),
		},
	}
	token, err := jwtclient.GenerateToken(&claims, secret.JwtSecret.MySecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateJWT(secret *ClientConnection, tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtclient.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.JwtSecret.MySecret), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*jwtclient.CustomClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}
