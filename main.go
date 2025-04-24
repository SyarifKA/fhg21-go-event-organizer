package main

import (
	"github.com/SyarifKA/fgh21-go-event-organizer/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/img/profile", "./img/profile")
	r.Static("/img/event", "./img/event")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	routers.RoutersCombine(r)

	// r.Run("0.0.0.0:8888")
	r.Run("localhost:8888")
}
