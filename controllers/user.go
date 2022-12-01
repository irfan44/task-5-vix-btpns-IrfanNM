package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/app"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/database"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/models"
)

func Login(c *gin.Context) {
	var user models.User
	var token string
	c.BindJSON(&user)
	db := database.InitDB()
	defer db.Close()
	if err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "unauthorized"})
		return
	}
	token, err := app.GenerateToken(user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": token})
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.Initialize()
	db := database.InitDB()
	defer db.Close()
	db.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User item created successfully!", "resourceId": user.ID})
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db := database.InitDB()
	defer db.Close()
	db.First(&user, id)
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db := database.InitDB()
	defer db.Close()
	db.First(&user, id)
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db := database.InitDB()
	defer db.Close()
	db.First(&user, id)
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
