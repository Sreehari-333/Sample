package controllers

import (
	model "Sample/Model"
	"Sample/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []model.Signup

	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch turfs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}
