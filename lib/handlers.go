package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOK(c *gin.Context, msg string, data any, pageInfo any) {
	c.JSON(http.StatusOK, Response{
		Success:  true,
		Message:  msg,
		PageInfo: pageInfo,
		Results:  data,
	})
}

func HandlerNotfound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: msg,
	})
}

func HandlerUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: msg,
	})
}

func HandlerBadReq(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: msg,
	})
}

func HandlerMaxFile(c *gin.Context, msg string) {
	c.JSON(http.StatusRequestEntityTooLarge, Response{
		Success: false,
		Message: msg,
	})
}
