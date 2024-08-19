package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventsRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllEvents)
	routerGroup.GET("/:id", controllers.DetailEvent)
	routerGroup.GET("/section/:id", controllers.FindSectionsByEventId)
	routerGroup.GET("/payment_method",middlewares.AuthMiddleware(), controllers.FindAllPaymentMethod)
	routerGroup.POST("",middlewares.AuthMiddleware(), controllers.ListCreateEvent)
	routerGroup.PATCH("/:id",middlewares.AuthMiddleware(), controllers.UpdateEvent)
	routerGroup.DELETE("/:id",middlewares.AuthMiddleware(), controllers.DeleteEvent)
}