package routes

import (
	"Sample/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.Group("/auth")
	{
		router.POST("/signup", controllers.Register)
		router.POST("/login", controllers.Login)
	}
}
