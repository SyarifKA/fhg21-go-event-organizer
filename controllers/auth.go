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

// type FormRegister struct {
// 	// Id int `json:"id"`
// 	FullName        string `json:"fullName" form:"fullName" db:"full_name"`
// 	Email           string `json:"email" form:"email" db:"email"`
// 	Password        string `json:"-" form:"password" db:"password"`
// 	ConfirmPassword string `json:"-" form:"confirmPassword" binding:"eqfield=password" db:"password"`
// }

func AuthRegister(ctx *gin.Context) {
	form := models.JoinProfile{}
	// form := FormRegister{}
	// var user models.User
	// var profile models.Profile

	err := ctx.Bind(&form)
	fmt.Println(form)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Register Failed",
			Results: form,
		})
		return
	}

	createProfile := models.CreateProfile(form)

	// user.Email = form.Email
	// user.Password = form.Password
	// profile.FullName = form.FullName
	// createUser := models.CreateUser(user)

	// userId := createUser.Id
	// profile.UserId = userId

	// createProfile := models.CreateProfile(profile)
	// createProfile.Email = form.Email
	// createProfile.FullName = form.FullName
	// createProfile.Password = form.Password

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

	if found == (models.User{}) {
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