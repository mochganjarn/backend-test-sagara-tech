package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/common/middleware"
	"github.com/mochganjarn/go-template-project/handler"
	"github.com/mochganjarn/go-template-project/service"
)

func productRoute(r *gin.RouterGroup, dependencies *service.ClientConnection) {
	v1 := r.Group("/v1/product")
	{
		v1.Use(middleware.ValidateToken(dependencies))
		v1.GET("/", handler.GetProduct(dependencies))
		v1.GET("/:id", handler.ShowProduct(dependencies))
		v1.POST("/", handler.CreateProduct(dependencies))
		v1.PUT("/:id", handler.UpdateProduct(dependencies))
		v1.DELETE("/:id", handler.DeleteProduct(dependencies))
	}
}
