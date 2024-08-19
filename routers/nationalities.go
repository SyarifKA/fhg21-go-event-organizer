package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func NationalityRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", controllers.ListAllNationalities)
}