package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func WishlistRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("/:id", controllers.CreateWishlist)
	routerGroup.GET("", controllers.ListAllWishlist)
	routerGroup.GET("/events", controllers.ListWishlistEvent)
}
