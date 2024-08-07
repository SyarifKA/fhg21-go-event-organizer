package controllers

import (
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {

	results := models.FindAllUsers(models.Data)
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all users",
		Results: results,
	})
}

func DetailUser(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := models.GetOneUser(models.Data, id)
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
    user := models.User{}

    ctx.Bind(&user)

    data := models.CreateUser(user)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Create User Success",
        Results: data,
    })
}

func DeleteUser(ctx *gin.Context){

}

func UpdateUser(c *gin.Context) {
    id, _:= strconv.Atoi(c.Param("id"))
    updatedData := models.User{}

    c.Bind(&updatedData)

    data := models.UpdateDataById(updatedData, id)

    if data.Id != 0 {
        c.JSON(http.StatusOK, lib.Response{
            Success: true,
            Message: "Update Data Success",
            Results: data,
        })
    } else {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "User Not Found",
        })
    }
}

func DeleteUserById(c *gin.Context) {
    id, _:= strconv.Atoi(c.Param("id"))

    data := models.DeleteUser(id)

    if data.Id != 0 {
        c.JSON(http.StatusOK, lib.Response{
            Success: true,
            Message: "Delete Data Success",
            Results: data,
        })
    } else {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Data Not Found",
        })
    }

}