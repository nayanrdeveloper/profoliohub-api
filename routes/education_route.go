package routes

import (
	"profoliohub-api/controllers"
	"profoliohub-api/middleware"
	"profoliohub-api/repositories"
	"profoliohub-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterEducationRoutes(apiRoutes *gin.RouterGroup, db *mongo.Database) {
	educationRepo := repositories.NewEducationRepository(db)
	edcuationService := services.NewEducationService(educationRepo)
	educationController := controllers.NewEducationController(edcuationService)

	educationRoutes := apiRoutes.Group("/education")
	educationRoutes.Use(middleware.AuthMiddleware())

	educationRoutes.GET("/", educationController.GetEducations)
	educationRoutes.POST("/", educationController.CreateEducation)
	educationRoutes.PUT("/:id", educationController.UpdateEducation)
	educationRoutes.DELETE("/:id", educationController.DeleteEducation)
}
