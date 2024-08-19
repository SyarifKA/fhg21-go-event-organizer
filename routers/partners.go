package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func PartnersRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", controllers.ListAllPartners)

}