package routes

import (
	"profoliohub-api/controllers"
	"profoliohub-api/repositories"
	"profoliohub-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAuthRoutes(apiRoutes *gin.RouterGroup, db *mongo.Database) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(userService)

	apiRoutes.POST("/register", authController.Register)
	apiRoutes.POST("login", authController.Login)
}
