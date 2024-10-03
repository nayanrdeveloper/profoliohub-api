package utils

import "github.com/gin-gonic/gin"


func BuildErrorResponse(message string, err string) gin.H {
	return gin.H{
		"status": "error",
		"message": message,
		"error": err,
	}
}

func BuildSuccessResponse(message string, data interface{}) gin.H  {
	return gin.H{
		"status": "success",
		"message": message,
		"data": data,
	}
}

func ResponseWithError(c *gin.Context, statusCode int, message string, err string)  {
	c.JSON(statusCode, BuildErrorResponse(message, err))
}

func ResponseWithSucess(c *gin.Context, statusCode int, message string, data interface{})  {
	c.JSON(statusCode, BuildSuccessResponse(message, data))
}