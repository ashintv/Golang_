package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

var users = []User{}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid data",
			})
			return
		}
		users = append(users, user)
		ctx.JSON(200, gin.H{
			"message": "user created successfull",
			"data":    user,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	router.Run()
}
