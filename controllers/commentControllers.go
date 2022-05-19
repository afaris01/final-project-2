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

func BuatComment(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == AppJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Content,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func AmbilComment(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	User := models.User{}
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	User.ID = userID
	Comment.UserID = userID

	err := db.Where("user_id = ?", userID).Find(&Comment).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error
	errPhoto := db.Where("user_id = ?", userID).Find(&Photo).Error

	if err != nil || errUser != nil || errPhoto != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Content,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
		"created_at": Comment.CreatedAt,
		"User": gin.H{
			"id":       User.ID,
			"email":    User.Email,
			"username": User.Username,
		},
		"Photo": gin.H{
			"id":        Photo.ID,
			"title":     Photo.Title,
			"caption":   Photo.Caption,
			"photo_url": Photo.PhotoURL,
			"user_id":   Photo.UserID,
		},
	})
}

func UbahComment(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if contentType == AppJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentId)
	Comment.UserID = uint(userId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Content: Comment.Content}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func HapusComment(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	Comment.ID = uint(commentId)
	Comment.UserID = userId

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}