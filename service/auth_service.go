package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	jwtclient "github.com/mochganjarn/go-template-project/external/jwt_client"
)

func JwtTokenGenerator(secret *ClientConnection, userID string) (string, error) {
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
