package controllers

import (
	"final-project-2/database"
	"final-project-2/helpers"
	"final-project-2/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

func BuatPhoto(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == AppJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func AmbilPhoto(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	User := models.User{}
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	Photo.UserID = userID

	err := db.Where("user_id = ?", userID).Find(&Photo).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error

	if err != nil || errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if Photo.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoURL,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
			"comments": gin.H{
				"id":    Comment.ID,
				"content": Comment.Content,
				"user_id": Comment.UserID,
				"created_at": Comment.CreatedAt,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "You don't have any photo",
		})
	}

}

func UbahPhoto(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if contentType == AppJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
}

func HapusPhoto(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	Photo.ID = uint(photoId)
	Photo.UserID = userId

	err := db.Model(&Photo).Where("id = ?", photoId).Delete(models.Photo{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
