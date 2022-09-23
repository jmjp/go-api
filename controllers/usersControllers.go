package controllers

import (
	"github.com/gin-gonic/gin"
	"un1versum.com/clutchers/bootstrap"
	dto "un1versum.com/clutchers/dtos/users"
	"un1versum.com/clutchers/models"
)

func FindUsers(c *gin.Context) {
	var user []models.User
	result := bootstrap.DB.Find(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"users": user,
	})
}

func CreateUser(c *gin.Context) {
	var body dto.CreateUserRequest
	c.Bind(&body)
	user := models.User{Username: body.Username}

	result := bootstrap.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func BodyCreateUser(c *gin.Context) {
	var body dto.CreateUserRequest
	c.BindJSON(body)
	c.JSON(200, gin.H{
		"user": body,
	})
}
