package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/data"
	"microservice/models"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, data.Users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range data.Users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = len(data.Users) + 1
	data.Users = append(data.Users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated models.User
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range data.Users {
		if user.ID == id {
			data.Users[i].Name = updated.Name
			data.Users[i].Email = updated.Email
			c.JSON(http.StatusOK, data.Users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range data.Users {
		if user.ID == id {
			data.Users = append(data.Users[:i], data.Users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "O‘chirildi"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
}
