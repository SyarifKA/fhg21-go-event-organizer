package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func CreateWishlist(ctx *gin.Context) {
	userId := ctx.GetInt("userId")

	eventId, _ := strconv.Atoi(ctx.Param("id"))

	dataWishlist := models.FindAllWishlist()
	fmt.Println(dataWishlist)

	for _, item := range dataWishlist {
		if item.UserId == userId && item.EventId == eventId {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: "Cannot add same wishlist",
				// Results: inputWishlist,
			})
			return
		}
	}

	inputWishlist := models.InputWishList(userId, eventId)

	// fmt.Println(inputWishlist.EventId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Add to wishlist success",
		Results: inputWishlist,
	})
}

func ListAllWishlist(ctx *gin.Context) {
	results := models.FindAllWishlist()

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all wishlist",
		Results: results,
	})
}

func ListWishlistEvent(ctx *gin.Context) {
	var results []models.Events
	for _, item := range models.FindAllWishlist() {
		results = append(results, models.FindOneEventById(item.EventId))
		// results = models.FindOneEventById(item.EventId)
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all events wishlist",
		Results: results,
	})
}
