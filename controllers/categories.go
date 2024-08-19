package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllCategories(ctx *gin.Context) {
	search := ctx.Query("search")
	limitParam := ctx.Query("limit")
	limit, _ := strconv.Atoi(limitParam)
	pageParam := ctx.Query("page")
	page, _ := strconv.Atoi(pageParam)

	if limit == 0 {
		limit = 7
	}

	if page == 0 {
		page = 1
	}

	// if page > 1 {
	// 	page = (page -1) * limit
	// }
	
	
	results, count := models.FindAllCategories(search, limit, page)
	totalPage := math.Ceil(float64(count)/float64(limit))

	if page > int(totalPage){
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Page not found",
			// PageInfo: pageInfo,
			// Results: results,
		})
		return
	}

	next := int(totalPage) - page
	prev := page - 1
	
	pageInfo := lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page: page,
		Limit: limit,
		Next: next,
		Prev: prev,
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Categories",
		PageInfo: pageInfo,
		Results: results,
	})
}

func DetailCategories(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categories := models.FindOneCategoriesById(id)
	if categories != (models.Categories{}){
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail categories",
			Results: categories,
		})
	}else{
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Categories not found",
		})	
	}
}

func CreateCategories(ctx *gin.Context) {
	newCategories := models.Categories{}

	if err := ctx.ShouldBind(&newCategories)
	err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	data := models.CreateCategories(newCategories)
	if data == (models.Categories{}) {
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

func UpdateCategories(c *gin.Context) {
    param := c.Param("id")
    id, _ := strconv.Atoi(param)
    // data := models.FindAllUsers()
    categories := models.Categories{}
    err := c.Bind(&categories)
    if err != nil {
        fmt.Println(err)
        return
    }
	categoriesUpdated := models.EditCategories(categories, id)

    // result := models.User{}
    // for _, v := range data {
    //     if v.Id == id {
    //         result = v
    //     }
    // }

    if categoriesUpdated.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "categories whit id " + param + " not found",
        })
        return
    }
    // models.EditUser(categories.Email, categories.Username, categories.Password, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "categories with id " + param + " Edit Success",
        Results: categoriesUpdated,
    })
}

func DeleteCategories(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneUserById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid categories Id",
		})
		return
	}
	err = models.DeleteCategories(id)
	fmt.Println(err)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Delete data category successfully",
		Results: dataUser,
	})
}