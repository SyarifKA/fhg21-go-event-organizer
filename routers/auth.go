package routers

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.AuthLogin)
	r.POST("/register", controllers.AuthRegister)
}
