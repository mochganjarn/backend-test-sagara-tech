package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/service"
)

func ValidateToken(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA)+1 {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Result": "Missing Token",
			})
			c.Abort()
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		status, err := service.ValidateJWT(appDependencies, tokenString)
		if err != nil {
			c.PureJSON(http.StatusUnauthorized, gin.H{
				"status": status,
				"result": "Invalid Token",
			})
			c.Abort()
		}
	}
}
