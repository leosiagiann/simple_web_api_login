package controllers

import (
	"Simple_Web_API_Login/database"
	"Simple_Web_API_Login/helpers"
	"Simple_Web_API_Login/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	user := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	// hash password
	user.Password = helpers.HashPassword(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	user := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password = user.Password

	// get user by email
	err := db.Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "invalid email",
		})
		return
	}

	// check password
	compare := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !compare {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "invalid password",
		})
		return
	}

	// generate token
	token := helpers.GenerateJwtToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
