package routers

import "github.com/gin-gonic/gin"

func RoutersCombine(r *gin.Engine) {
		UserRouter(r.Group("/users"))
		AuthRouter(r.Group("/auth"))
		EventsRouter(r.Group("/events"))
		CategoriesRouter(r.Group("/categories"))
		TransactionRouter(r.Group("/transactions"))
		ProfileRouter(r.Group("/profile"))
		PartnersRouter(r.Group("/partners"))
		NationalityRouter(r.Group("/nationalities"))
}