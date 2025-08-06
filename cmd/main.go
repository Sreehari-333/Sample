package main

import (
	"Sample/db"
	"Sample/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDb()

	router := gin.Default()
	routes.AuthRoutes(router)
	// routes.PublicRoutes(router)

	router.Run(":8080")
}
