package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func CreateWishlist(ctx *gin.Context) {
	userId := ctx.GetInt("userId")

	eventId, _ := strconv.Atoi(ctx.Param("id"))

	dataWishlist := repository.FindAllWishlist()
	fmt.Println(dataWishlist)

	for _, item := range dataWishlist {
		if item.UserId == userId && item.EventId == eventId {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: "Cannot add same wishlist",
			})
			return
		}
	}

	inputWishlist := repository.InputWishList(userId, eventId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Add to wishlist success",
		Results: inputWishlist,
	})
}

func ListAllWishlist(ctx *gin.Context) {
	results := repository.FindAllWishlist()

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all wishlist",
		Results: results,
	})
}

func ListWishlistEvent(ctx *gin.Context) {
	var results []dtos.Events
	for _, item := range repository.FindAllWishlist() {
		results = append(results, repository.FindOneEventById(item.EventId))
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all events wishlist",
		Results: results,
	})
}
