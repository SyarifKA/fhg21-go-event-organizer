package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllEvents)
	routerGroup.GET("/:id", controllers.DetailEvent)
	routerGroup.GET("/section/:id", controllers.FindSectionsByEventId)
	routerGroup.GET("/payment_method", controllers.FindAllPaymentMethod)
	routerGroup.POST("", controllers.ListCreateEvent)
	routerGroup.PATCH("/:id", controllers.UpdateEvent)
	routerGroup.DELETE("/:id", controllers.DeleteEvent)
}