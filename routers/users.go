package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllUser)
	routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("", controllers.CreateUser)
	routerGroup.PATCH("/:id", controllers.UpdateUser)
	routerGroup.DELETE("/:id", controllers.DeleteUser)
}