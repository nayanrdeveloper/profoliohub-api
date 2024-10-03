package main

import (
	"log"
	"profoliohub-api/config"
	"profoliohub-api/database"
	"profoliohub-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := database.ConnectDB()

	r := gin.Default()
	routes.SetUpRoute(r,db)
	log.Fatal(r.Run(":8080"))
}