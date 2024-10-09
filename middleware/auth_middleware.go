package middleware

import (
	"net/http"
	"profoliohub-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthorized", "Missing Token"))
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		user_id, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthorized", "Invalid Token"))
			c.Abort()
			return
		}

		userId, _ := primitive.ObjectIDFromHex(user_id)
		c.Set("user_id", userId)
		c.Next()
	}
}
