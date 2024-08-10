package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func EventsRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllEvents)
	// routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("", controllers.CreateEvent)
	// routerGroup.PATCH("/:id", controllers.UpdateUser)
	// routerGroup.DELETE("/:id", controllers.DeleteUser)
}