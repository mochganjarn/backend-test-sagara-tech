package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/handler"
	"github.com/mochganjarn/go-template-project/service"
)

func userRoute(r *gin.RouterGroup, dependencies *service.ClientConnection) {
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handler.HelloWorld(dependencies))
		v1.GET("/generate-token", handler.ClaimToken(dependencies))
	}
}
