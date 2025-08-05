package controllers

import (
	model "Sample/Model"
	"Sample/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup

func Register(c *gin.Context) {
	var data struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "inavlid input"})
		return
	}

	newuser := model.Signup{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := db.DB.Create(&newuser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

}
