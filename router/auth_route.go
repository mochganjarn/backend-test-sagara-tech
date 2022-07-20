package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/handler"
	"github.com/mochganjarn/go-template-project/service"
)

func authRoute(r *gin.RouterGroup, dependencies *service.ClientConnection) {
	v1 := r.Group("/v1")
	{
		v1.GET("/login", handler.Login(dependencies))
		v1.POST("/register", handler.Register(dependencies))
	}

}
