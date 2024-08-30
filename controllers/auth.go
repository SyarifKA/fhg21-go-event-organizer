package controllers

import (
	"fmt"
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

// func AuthRegister(ctx *gin.Context) {
// 	form := dtos.JoinProfile{}
// 	var user dtos.User

// 	err := ctx.Bind(&form)
// 	fmt.Println(form)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Register Failed",
// 		})
// 		return
// 	}

// 	dataUser := repository.FindOneUserByEmail(form.Email)

// 	if dataUser.Email == form.Email {
// 		ctx.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Email already exist",
// 		})
// 		return
// 	}

// 	repository.CreateProfile(form)

// 	user.Email = form.Email
// 	user.Password = form.Password
// 	createUser := repository.CreateUser(user)

// 	ctx.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "Register Success",
// 		Results: createUser,
// 	})
// }

func AuthRegister(c *gin.Context) {
	formRegister := dtos.RegisterForm{}
	err := c.Bind(&formRegister)
	fmt.Println(formRegister)
	if err != nil {
		lib.HandlerBadReq(c, "format invalid")
		return
	}

	user, err := repository.CreateUser(models.Users{
		Email:    formRegister.Email,
		Password: formRegister.Password,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	profile, err := repository.CreateProfile(models.Profile{
		FullName: formRegister.FullName,
		UserId:   user.Id,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Register success", profile, nil)
}

func AuthLogin(ctx *gin.Context) {
	var user dtos.User
	ctx.Bind(&user)

	found := repository.FindOneUserByEmail(user.Email)
	// fmt.Println(found)

	if found == (dtos.User{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password1",
		})
		return
	}

	if found.Email != user.Email && found.Password != user.Password {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password2",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)
	// fmt.Println(user.Password)
	// fmt.Println(found.Password)
	// fmt.Println(found.Id)
	// fmt.Println(isVerified)
	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		lib.HandlerOK(ctx, "Login success", dtos.Token{Token: JWToken}, nil)
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password3",
		})
	}
}
