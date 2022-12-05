package routes

import (
	"cake_store/controller"

	"github.com/gin-gonic/gin"
)

func AddCakeRoutes(router *gin.Engine) {
	routerV1 := router.Group("/v1")

	{
		routerV1.GET("/cakes/:id", controller.GetCakeDetail)
		routerV1.PATCH("/cakes/:id", controller.UpdateCake)
		routerV1.DELETE("/cakes/:id", controller.DeleteCake)
		routerV1.GET("/cakes", controller.GetCakes)
		routerV1.POST("/cakes", controller.CreateCake)
	}
}
