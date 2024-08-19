package controllers

import (
	"fmt"
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func DataProfile(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	fmt.Println(userId)
	dataProfile := models.FindProfileByUserId(userId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail Profile",
		Results: dataProfile,
	})
}

func UpdateProfile(ctx *gin.Context) {

	id := ctx.GetInt("userId")
    profile := models.Profiles{}
    err := ctx.Bind(&profile)
    if err != nil {
        fmt.Println(err)
        return
    }
	profileUpdated := models.EditProfile(profile, id)

    if profileUpdated.Id == 0 {
        ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Profile not found",
        })
        return
    }

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Profile updated",
        Results: profileUpdated,
    })
}