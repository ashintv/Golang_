package models

type User struct {
	Name string `json:"name" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

var Users = []User{}

