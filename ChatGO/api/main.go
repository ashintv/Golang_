package main

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	usser_api := router.Group("/user")
	{
		usser_api.POST("/", controllers.CreateUser)
		usser_api.GET("/", controllers.GetAllUsers)
	}

	router.Run()
}
