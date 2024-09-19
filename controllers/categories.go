package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
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

	results, count := repository.FindAllCategories(search, limit, page)
	totalPage := math.Ceil(float64(count) / float64(limit))

	if page > int(totalPage) {
		lib.HandlerBadReq(ctx, "Page not found")
		return
	}

	next := int(totalPage) - page
	prev := page - 1

	pageInfo := lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	lib.HandlerOK(ctx, "List All Categories", results, pageInfo)
}

func DetailCategories(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categories := repository.FindOneCategoriesById(id)
	if categories != (dtos.Categories{}) {
		lib.HandlerOK(ctx, "Detail Categories", categories, nil)
	} else {
		lib.HandlerNotfound(ctx, "Categories not found")
	}
}

func CreateCategories(ctx *gin.Context) {
	newCategories := dtos.Categories{}

	if err := ctx.ShouldBind(&newCategories); err != nil {
		lib.HandlerBadReq(ctx, "Invalid input data")
		return
	}

	data := repository.CreateCategories(newCategories)
	if data == (dtos.Categories{}) {
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
	categories := dtos.Categories{}
	err := c.Bind(&categories)
	if err != nil {
		fmt.Println(err)
		return
	}
	categoriesUpdated := repository.EditCategories(categories, id)

	if categoriesUpdated.Id == 0 {
		c.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "categories whit id " + param + " not found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "categories with id " + param + " Edit Success",
		Results: categoriesUpdated,
	})
}

func DeleteCategories(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataUser := repository.FindOneUserById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid categories Id",
		})
		return
	}
	err = repository.DeleteCategories(id)
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
		Message: "Delete data category successfully",
		Results: dataUser,
	})
}

func FindEventsByCategoryId(ctx *gin.Context) {
	form := dtos.GetEventCategories{}

	err := ctx.Bind(&form)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(form.CategoryId)

	result, err := repository.FindEventByCategoryId(form.CategoryId)

	fmt.Println(result)

	if err != nil {
		lib.HandlerBadReq(ctx, "failed get event")
		return
	}

	lib.HandlerOK(ctx, "Success get event by Id", result, nil)
}
