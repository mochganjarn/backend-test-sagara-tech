package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/service"
)

func Init(dependencies *service.ClientConnection) {
	r := gin.Default()
	api := r.Group("/api")
	authRoute(api, dependencies)
	productRoute(api, dependencies)
	httpPort := dependencies.AppPort
	if httpPort == "" {
		httpPort = "8080"
	}
	r.Run(":" + httpPort)
}
