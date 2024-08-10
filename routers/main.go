package routers

import "github.com/gin-gonic/gin"

func RoutersCombine(r *gin.Engine) {
		UserRouter(r.Group("/users"))
		AuthRouter(r.Group("/auth"))
		EventsRouter(r.Group("/events"))
}