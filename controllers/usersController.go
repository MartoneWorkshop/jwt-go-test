package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/MartoneWorkshop/jwt-go-test/initializers"
	"github.com/MartoneWorkshop/jwt-go-test/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struc {
		Email 	 string
		Password string
	}
	if c.Bind(&body) != nill {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nill {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	f err != nill {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H)
}