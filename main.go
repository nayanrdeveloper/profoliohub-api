package main

import (
	"log"
	"profoliohub-api/config"
	"profoliohub-api/database"
	"profoliohub-api/middleware"
	"profoliohub-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := database.ConnectDB()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.SetUpRoutes(router, db)
	log.Fatal(router.Run(":8080"))
}
