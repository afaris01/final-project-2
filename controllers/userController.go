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

type Input struct {
	Age      uint   `json:"age" bit inding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var (
	AppJson  = "application/json"
	User     models.User
	regInput Input
	upInput  Input
)

func UserRegister(c *gin.Context) {
	db := database.AmbilDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	if contentType == AppJson {
		err := c.ShouldBindJSON(&regInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		c.ShouldBind(&regInput)
	}

	User := models.User{Age: regInput.Age, Email: regInput.Email, Username: regInput.Username, Password: regInput.Password}
	database.DB.Create(&User)

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
	User := models.User{}
	_, _ = db, contentType

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
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	paramId, _ := strconv.Atoi(c.Param("id"))
	userId := uint(userData["id"].(float64))
	User.ID = userId

	if contentType == AppJson {
		err := database.DB.Where("id = ?", paramId).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Record Not Found ",
				"message": err.Error(),
			})
			return
		}
	}

	if err := c.ShouldBindJSON(&upInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	upInput.Password = helpers.HashPass(upInput.Password)
	update := models.User{Age: upInput.Age, Email: upInput.Email, Username: upInput.Username, Password: upInput.Password}

	if err := database.DB.Model(&User).Updates(update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	err := db.Model(&User).Where("id = ?", User.ID).Delete(models.User{}).Error

	if err != nil && userID != User.ID {
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
