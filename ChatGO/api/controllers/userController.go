package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid data",
		})
		return
	}
	models.Users = append(models.Users, user)
	c.JSON(200 , "user created succes fully")
}

func GetAllUsers(c *gin.Context){
	c.JSON(200 , gin.H{
		"users":models.Users,
	})
}