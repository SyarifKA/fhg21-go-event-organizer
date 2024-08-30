package controllers

import (
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func ListAllLocations(ctx *gin.Context) {
	results := repository.FindAllLocation()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Locations",
		Results: results,
	})
}
