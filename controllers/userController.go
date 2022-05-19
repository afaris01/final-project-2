package controllers

import (
	"final-project-2/database"
	"final-project-2/helpers"
	"final-project-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	AppJson = "application/json"

func UserRegister(c *gin.Context) {
	db := database.AmbilDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	var (
		// tambahUser user
		User models.User
	)

	if contentType == AppJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := c.ShouldBindJSON(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"age":        User.Age,
		"email":      User.Email,
		"password":   User.Password,
		"username":   User.Username,
		"created_at": User.CreatedAt,
	})
}

func UserLogin(c *gin.Context) {
	db := database.AmbilDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == AppJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UbahUser(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	paramId, _ := strconv.Atoi(c.Param("userId"))
	userId := uint(userData["id"].(float64))

	if contentType == AppJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = userId

	err := db.Model(&User).Where("id = ?", paramId).Updates(models.User{Username: User.Username, Email: User.Email, Password: User.Password, Age: User.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"age":        User.Age,
		"email":      User.Email,
		"password":   User.Password,
		"username":   User.Username,
		"updated_at": User.UpdatedAt,
	})
}

func HapusUser(c *gin.Context) {
	db := database.AmbilDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	User.ID = userID

	err := db.Model(&User).Where("id = ?", userID).Delete(models.User{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
