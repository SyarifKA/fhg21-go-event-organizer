package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

type Response struct{
	Success bool `json:"succes"`
	Message string `json:"message"`
	Results interface{} `json:"results" binding:"omitempty"`
}

type Auth struct{
	Email string
	Password string
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	data := []User{}
	r.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "List all users",
			Results: data,
		})
	})
	r.POST("/users", func(c *gin.Context){
		user := User{}

		c.Bind(&user)

		conv := c.Bind(&user)

		dataEmail := user.Email
		// fmt.Println(dataEmail)

		// user.Id = len(data)+1

		result := 0

		for _, x := range data{
			result = x.Id
		}

		user.Id = result +1
		
		if conv == nil{
			compar := true
			for compar{
				for _, item := range data{
					// result = item.Id
					// fmt.Println(item)
					if dataEmail == item.Email{
					// if user == item{
						c.JSON(http.StatusBadRequest, Response{
							Success: false,
							Message: "Email already exist",
							// Results: user,
						})
						return
					}
				}
				compar = false
			}
			data = append(data, user)
			c.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Create user sucess",
				Results: user,
			})
		}else{
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Account does not match the criteria",
				// Results: user,
			})
		}
	})
	r.POST("/auth/login", func(c *gin.Context){
		user := Auth{}

		c.Bind(&user)

		email := user.Email
		// password := user.Password

		searchData := true
			for searchData {
				for i := 0; i<len(data); i++{
					dataEmail := data[i].Email
					if email == dataEmail{
						c.JSON(http.StatusOK, Response{
							Success: true,
							Message: "Login success",
							// Results: data,
						})
						return
					}
				}
				searchData = false
			}
			c.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "wrong email or password",
				// Results: data,
			})
	})
	r.GET("/users/:id", func(c *gin.Context){
		userId, _ := strconv.Atoi(c.Param("id"))

		selected := -1

		for index, item := range data{
			if item.Id == userId{
				selected = index
			}
		}

		if selected != -1 {
			c.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Data user",
					Results: data[selected],
				})
		}else{
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "User not found",
				// Results: data[selected],
			})
		}

		// searchId := true
		// 	for searchId{
		// 		for i := 0; i<len(data); i++{
		// 			dataId := data[i].Id
		// 			if user == dataId{
		// 				c.JSON(http.StatusOK, Response{
		// 					Success: true,
		// 					Message: "Data user",
		// 					Results: data[user],
		// 				})
		// 				return
		// 			}
		// 		}
		// 		searchId = false
		// 	}
		// 	c.JSON(http.StatusUnauthorized, Response{
		// 		Success: false,
		// 		Message: "Data not found",
		// 		// Results: data,
		// 	})
	})
	// r.DELETE("/users/:id", func(c *gin.Context){
	// 	userId, _ := strconv.Atoi(c.Param("id"))

	// 	user := userId

	// 	searchId := true
	// 		for searchId{
	// 			for i := 0; i<len(data); i++{
	// 				dataId := data[i].Id
	// 				if user == dataId{
	// 					c.JSON(http.StatusOK, Response{
	// 						Success: true,
	// 						Message: "Data user",
	// 						Results: data,
	// 					})
	// 					return
	// 				}
	// 			}
	// 			searchId = false
	// 		}
	// 		c.JSON(http.StatusUnauthorized, Response{
	// 			Success: false,
	// 			Message: "Data not found",
	// 			// Results: data,
	// 		})
	// })
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		userId, _ := strconv.Atoi(ctx.Param("id"))

		selected := -1

		for index, item := range data{
			if item.Id == userId{
				selected = index
			}
		}

		if selected != -1 {
			form := User{}
			ctx.Bind(&form)

			updateData := ctx.Bind(&form)

			if updateData == nil{
				data[selected].Name = form.Name
				data[selected].Email = form.Email
				data[selected].Password = form.Password
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Update User Success",
					Results: data[selected],
				})
			}else{
				ctx.JSON(http.StatusNotFound, Response{
					Success: false,
					Message: "Data does not match the criteria",
				})
			}
		}else{
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			})
		}
	})
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		userId, _ := strconv.Atoi(ctx.Param("id"))

		selected := -1

		for index, item := range data{
			if item.Id == userId{
				selected = index
			}
		}

		if selected != -1 {
			dataSelected := data[selected]
			sliceLeft := data[0:selected]
			sliceRight := data[selected+1:]
			data = sliceLeft
			data = append(data, sliceRight...)
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Update User Success",
				Results: dataSelected,
			})
		}else{
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			})
		}
	})
	r.Run("localhost:8888")
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}