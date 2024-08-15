package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func TransactionRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	// routerGroup.GET("", controllers.ListAllCategories)
	routerGroup.GET("", controllers.FindTransactionByUserId)
	routerGroup.POST("", controllers.CreateTransaction)
	// routerGroup.PATCH("/:id", controllers.UpdateCategories)
	// routerGroup.DELETE("/:id", controllers.DeleteCategories)
}