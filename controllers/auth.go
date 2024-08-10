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

func AuthRegister(ctx *gin.Context){
	// var form FormRegister
	// var user models.User
	// var profile models.Profile

	form := models.FormRegister{}

	err := ctx.Bind(&form)
	// fmt.Println(form)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to Register",
		})
		return
	}

	fmt.Println(form)

	data := models.RegisterUser(form)
	if data == (models.FormRegister{}) {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to Register",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Register successfully",
		Results: data,
	})

	// id, _ := strconv.Atoi(ctx.Param("id"))
	// user := models.FindProfileByUserId(id)
	// if user != (models.Profile{}){
	// 	ctx.JSON(http.StatusOK, lib.Response{
	// 		Success: true,
	// 		Message: "Detail user",
	// 		Results: user,
	// 	})
	// }else{
	// 	ctx.JSON(http.StatusNotFound, lib.Response{
	// 		Success: false,
	// 		Message: "User not found",
	// 	})	
	// }
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