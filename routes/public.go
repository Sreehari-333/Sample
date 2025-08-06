package routes

import (
	"Sample/controllers"

	"github.com/gin-gonic/gin"
)

func Public(router *gin.Engine) {

	router.GET("/users", controllers.GetAllUsers)

}
