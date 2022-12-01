package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/models"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/types"
	"github.com/jinzhu/gorm"
)

func GetPhoto(c *gin.Context) {
	photos := []models.Photo{}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Debug().Model(&models.Photo{}).Limit(100).Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": "photo not found", "data": nil})
		return
	}

	if len(photos) > 0 {
		for i := range photos {
			user := models.User{}
			err := db.Model(&models.User{}).Where("id = ?", photos[i].UserId).Take(&user).Error

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
				return
			}

			photos[i].Author = types.Author{
				ID: user.ID, Username: user.Username, Email: user.Email,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "T", "message": "success", "data": photos})
}

// create photo with jwt with user id, author, id, username, email
func CreatePhoto(c *gin.Context) {
	photo := models.Photo{}

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	photo.UserId = c.MustGet("user_id").(string)
	photo.Author = types.Author{
		ID:       c.MustGet("user_id").(string),
		Username: c.MustGet("username").(string),
		Email:    c.MustGet("email").(string),
	}

	if err := db.Debug().Model(&models.Photo{}).Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "T", "message": "success", "data": photo})
}

// update photo with jwt auth
func UpdatePhoto(c *gin.Context) {
	photo := models.Photo{}

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	if err := db.Debug().Model(&models.Photo{}).Where("id = ?", c.Param("id")).Updates(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "T", "message": "success", "data": photo})
}

// delete photo with jwt auth
func DeletePhoto(c *gin.Context) {
	photo := models.Photo{}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Debug().Model(&models.Photo{}).Where("id = ?", c.Param("id")).Take(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	if err := db.Debug().Model(&models.Photo{}).Where("id = ?", c.Param("id")).Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "F", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "T", "message": "success", "data": photo})
}
