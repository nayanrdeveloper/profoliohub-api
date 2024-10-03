package controllers

import (
	"net/http"
	"profoliohub-api/models"
	"profoliohub-api/services"
	"profoliohub-api/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (ctrl *AuthController) Register(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Invalid data", err.Error()))
        return
    }

	if err := ctrl.authService.RegisterUser(user); err != nil{
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Registration successful",nil))
}

func (ctrl *AuthController) Login(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusOK, utils.BuildErrorResponse("Invalid Data", err.Error()))
		return
	}

	token, err :=ctrl.authService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Authentication failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Login Successful", gin.H{"token": token}))

}