package controllers

import (
	"fmt"
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct{
	JWToken string `json:"token"`
}

func AuthRegister(ctx *gin.Context) {
	form := models.JoinProfile{}
	var user models.User
	// var profile models.Profile

	err := ctx.Bind(&form)
	fmt.Println(form)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Register Failed",
		})
		return
	}

	dataUser := models.FindOneUserByEmail(form.Email)
	fmt.Println(dataUser)

	if dataUser.Email == form.Email{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Email already exist",
		})
		return
	}

	createProfile := models.CreateProfile(form)

	user.Email = form.Email
	user.Password = form.Password
	models.CreateUser(user)


	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Register Success",
		Results: createProfile,
	})
}

func AuthLogin(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)

	found := models.FindOneUserByEmail(user.Email)
	fmt.Println(found)

	if found == (models.User{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}

	if found.Email != user.Email && found.Password != user.Password{
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified{
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Login success",
			Results: Token{
				JWToken,
			},
		})
	}else{
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password",
		})
	}
}