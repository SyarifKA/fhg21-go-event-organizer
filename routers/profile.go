package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/SyarifKA/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfileRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.DataProfile)
	routerGroup.GET("/login", controllers.FindProfileByUserId)
	routerGroup.PATCH("", controllers.UpdateProfile)
	routerGroup.PATCH("/img", controllers.UploadProfileImage)
}
