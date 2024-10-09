package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRoutes(router *gin.Engine, db *mongo.Database) {
	apiRoutes := router.Group("/api")

	RegisterAuthRoutes(apiRoutes, db)
	RegisterEducationRoutes(apiRoutes, db)
}
