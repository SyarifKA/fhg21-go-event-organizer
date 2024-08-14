package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoriesRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllCategories)
	routerGroup.GET("/:id", controllers.DetailCategories)
	routerGroup.POST("", controllers.CreateCategories)
	routerGroup.PATCH("/:id", controllers.UpdateCategories)
	routerGroup.DELETE("/:id", controllers.DeleteCategories)
}