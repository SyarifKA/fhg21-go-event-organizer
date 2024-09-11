package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func ListAllUser(ctx *gin.Context) {
	search := ctx.Query("search")
	limitParam := ctx.Query("limit")
	limit, _ := strconv.Atoi(limitParam)
	pageParam := ctx.Query("page")
	page, _ := strconv.Atoi(pageParam)

	if limit == 0 {
		limit = 3
	}

	if page == 0 {
		page = 1
	}

	// if page > 1 {
	// 	page = (page -1) * limit
	// }

	// totalData := models.TotalData(search)
	results, count := repository.FindAllUsers(search, limit, page)
	totalPage := math.Ceil(float64(count) / float64(limit))

	next := int(totalPage) - page
	prev := page - 1

	fmt.Println(results)
	fmt.Println(count)

	pageInfo := lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success:  true,
		Message:  "List All Users",
		PageInfo: pageInfo,
		Results:  results,
	})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := repository.FindOneUserById(id)
	if user != (dtos.User{}) {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail user",
			Results: user,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
	}
}

// func CreateUser(ctx *gin.Context) {
// 	newUser := dtos.User{}

// 	if err := ctx.ShouldBind(&newUser); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "invalid input data",
// 		})
// 		return
// 	}

// 	data, err := repository.CreateUser(newUser)
// 	if data == (dtos.User{}) {
// 		ctx.JSON(http.StatusBadRequest, lib.Response{
// 			Success: false,
// 			Message: "Failed to create user",
// 		})
// 		return
// 	}

// 	lib.HandlerOK(ctx, "User created successfully", data, nil)
// }

func CreateUser(c *gin.Context) {
	formUser := dtos.FormUser{}
	err := c.Bind(&formUser)

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	data, err := repository.CreateUser(models.Users{
		Email:    formUser.Email,
		Password: formUser.Password,
	})

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Create user success", data, nil)
}

func UpdateUser(c *gin.Context) {
	param := c.Param("id")
	// id, _ := strconv.Atoi(param)
	// data := models.FindAllUsers()
	id, _ := strconv.Atoi(param)
	user := dtos.Profiles{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	userUpdated, err := repository.EditUser(models.UserUpdate{
		Email:    *user.Email,
		Username: user.Username,
	}, id)
	fmt.Println(userUpdated)
	if err != nil {
		fmt.Println(err)
		return
	}

	if userUpdated.Id == 0 {
		c.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "user with id " + param + " not found",
		})
		return
	}
	// models.EditUser(user.Email, user.Username, user.Password, param)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "user with id " + param + " Edit Success",
		Results: userUpdated,
	})
}

func UpdateUserPassword(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	user := dtos.Password{}

	err := ctx.Bind(&user)
	if err != nil {
		lib.HandlerBadReq(ctx, "Failed to change password")
		return
	}

	result, err := repository.EditPassword(models.UpdatePassword{
		OldPassword:     user.OldPassword,
		NewPassword:     user.NewPassword,
		ConfirmPassword: user.ConfirmPassword,
	}, id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Update user password success",
		Results: result,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := repository.FindOneUserById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid user Id",
		})
		return
	}
	err = repository.DeleteUser(id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Delete data user successfully",
		Results: dataUser,
	})
}
