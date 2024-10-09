package controllers

import (
	"net/http"
	"profoliohub-api/models"
	"profoliohub-api/services"
	"profoliohub-api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EducationController struct {
	educationService *services.EducationService
}

func NewEducationController(educationService *services.EducationService) *EducationController {
	return &EducationController{
		educationService: educationService,
	}
}

func (ctrl *EducationController) CreateEducation(c *gin.Context) {
	var education models.Education
	userId := c.MustGet("user_id").(primitive.ObjectID)

	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Invalid Data", err.Error()))
		return
	}

	if err := ctrl.educationService.CreateEducation(userId, education); err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("Error creating education", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Education Created", nil))
}

func (ctrl *EducationController) GetEducations(c *gin.Context) {
	userId := c.MustGet("user_id").(primitive.ObjectID)

	education, err := ctrl.educationService.GetUserEducation(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("Error fetching educations", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Edcuation retrevied", education))
}

func (ctrl *EducationController) UpdateEducation(c *gin.Context) {
	var education models.Education
	user_id := c.MustGet("user_id").(primitive.ObjectID)

	// Bind request data to education object
	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Invalid Data", err.Error()))
		return
	}

	// Ensure that the UserId is correctly set
	education.UserId = user_id

	// Get the education ID from URL parameters
	educationId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Invalid education Id", err.Error()))
		return
	}

	education.ID = educationId

	// Call the service to update the education record
	if err := ctrl.educationService.UpdateEducation(user_id, education); err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("updating education", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Education Updated", nil))
}

func (ctrl *EducationController) DeleteEducation(c *gin.Context) {
	userID := c.MustGet("user_id").(primitive.ObjectID)
	educationId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Invalid education ID", err.Error()))
		return
	}

	if err := ctrl.educationService.DeleteEducatationById(userID, educationId); err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("Error deleting education", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse("Education deleted", nil))
}
