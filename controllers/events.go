package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

func DetailEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	eventDetail := models.FindOneEventById(id)
	fmt.Println(eventDetail)
	if eventDetail != (models.Events{}){
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail event",
			Results: eventDetail,
		})
	}else{
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
	}
}

func ListCreateEvent(ctx *gin.Context) {
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
	// fmt.Println(data)
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

func DeleteEvent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataEvent := models.FindOneEventById(id)
	// fmt.Println(dataEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid event Id",
		})
		return
	}
	err = models.DeleteEvent(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Delete data event successfully",
		Results: dataEvent,
	})
}

func UpdateEvent(c *gin.Context) {
    param := c.Param("id")
    id, _ := strconv.Atoi(param)
    // data := models.FindAllEvents()
	// dataEvent := models.FindOneEventById(id)

    event := models.Events{}

    err := c.Bind(&event)
    if err != nil {
        fmt.Println(err)
        return
    }
	dataUpdated := models.EditEvent(event, id)

    // result := models.Events{}
    // for _, v := range data {
    //     if v.Id == id {
    //         result = v
    //     }
    // }

    if dataUpdated.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "event with id " + param + " not found",
        })
        return
    }

    // eventDetail := models.EditEvent(event, id)
    // models.EditEvent(event.Image, event.Title, event.Date, event.Description, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "event with id " + param + " Edit Success",
        Results: dataUpdated,
    })
}