package controllers

import (
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllNationalities(ctx *gin.Context) {
	results := models.FindAllNationality()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Nationalities",
		Results: results,
	})
}