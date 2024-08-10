package controllers

import (
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllEvents(ctx *gin.Context) {
	results := models.FindAllEvents()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Events",
		Results: results,
	})
}

func CreateEvent(ctx *gin.Context) {
	newEvent := models.Events{}

	if err := ctx.ShouldBind(&newEvent)
	err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	data := models.CreateEvent(newEvent)
	if data == (models.Events{}) {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to create event",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Event created successfully",
		Results: data,
	})
}