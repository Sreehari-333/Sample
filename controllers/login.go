package controllers

import (
	model "Sample/Model"
	"Sample/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var user model.Signup

	err := db.DB.Where("email = ?", input.Email).First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "User login successfull",
		"User":    user,
	})
}
