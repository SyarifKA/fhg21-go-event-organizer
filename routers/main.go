package routers

import "github.com/gin-gonic/gin"

func RoutersCombine(r *gin.Engine) {
		UserRouter(r.Group("/users"))
}