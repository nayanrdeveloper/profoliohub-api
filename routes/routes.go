package routes

import (
	"profoliohub-api/controllers"
	"profoliohub-api/repositories"
	"profoliohub-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRoute(router *gin.Engine, db *mongo.Database) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}