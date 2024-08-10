package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

// func ListAllUsers(ctx *gin.Context) {
// 	results := models.FindAllUsers(models.Data)
// 	ctx.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "List all users",
// 		Results: results,
// 	})
// }

// func DetailUser(ctx *gin.Context){
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	user := models.GetOneUser(models.Data, id)
// 	if user != (models.User{}){
// 		ctx.JSON(http.StatusOK, lib.Response{
// 			Success: true,
// 			Message: "Detail user",
// 			Results: user,
// 		})
// 	}else{
// 		ctx.JSON(http.StatusNotFound, lib.Response{
// 			Success: false,
// 			Message: "User not found",
// 		})
// 	}
// }

// func CreateUser(ctx *gin.Context) {
//     user := models.User{}
//     ctx.Bind(&user)
//     data := models.CreateUser(user)
//     ctx.JSON(http.StatusOK, lib.Response{
//         Success: true,
//         Message: "Create User Success",
//         Results: data,
//     })
// }

// func UpdateUser(c *gin.Context) {
//     id, _:= strconv.Atoi(c.Param("id"))
//     updatedData := models.User{}
//     c.Bind(&updatedData)
//     data := models.UpdateDataById(updatedData, id)
//     if data.Id != 0 {
//         c.JSON(http.StatusOK, lib.Response{
//             Success: true,
//             Message: "Update Data Success",
//             Results: data,
//         })
//     } else {
//         c.JSON(http.StatusNotFound, lib.Response{
//             Success: false,
//             Message: "User Not Found",
//         })
//     }
// }

// func DeleteUserById(c *gin.Context) {
//     id, _:= strconv.Atoi(c.Param("id"))
//     data := models.DeleteUser(id)
//     if data.Id != 0 {
//         c.JSON(http.StatusOK, lib.Response{
//             Success: true,
//             Message: "Delete Data Success",
//             Results: data,
//         })
//     } else {
//         c.JSON(http.StatusNotFound, lib.Response{
//             Success: false,
//             Message: "Data Not Found",
//         })
//     }
// }

// Controllers users
func ListAllUser(ctx *gin.Context) {
	results := models.FindAllUsers()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Users",
		Results: results,
	})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := models.FindOneUserById(id)
	if user != (models.User{}){
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail user",
			Results: user,
		})
	}else{
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})	
	}
}

func CreateUser(ctx *gin.Context) {
	newUser := models.User{}

	if err := ctx.ShouldBind(&newUser)
	err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	data := models.CreateUser(newUser)
	if data == (models.User{}) {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User created successfully",
		Results: data,
	})
}

func UpdateUser(c *gin.Context) {
    param := c.Param("id")
    id, _ := strconv.Atoi(param)
    data := models.FindAllUsers()

    user := models.User{}
    err := c.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }
    result := models.User{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "user whit id " + param + " not found",
        })
        return
    }
    models.EditUser(user.Email, user.Username, user.Password, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user with id " + param + " Edit Success",
        Results: user,
    })
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneUserById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid user Id",
		})
		return
	}
	err = models.DeleteUser(id)
	if err != nil{
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