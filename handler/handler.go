package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/service"
)

func HelloWorld(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	}
}

func ClaimToken(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.JwtTokenGenerator(appDependencies, "iduser")
		if err != nil {
			c.PureJSON(400, gin.H{
				"Result": "Failed Generate Token",
			})
			return
		}
		c.PureJSON(200, gin.H{
			"Result": token,
		})
	}
}
