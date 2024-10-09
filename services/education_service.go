package services

import (
	"errors"
	"profoliohub-api/models"
	"profoliohub-api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EducationService struct {
	educationRepository *repositories.EducationRepository
}

func NewEducationService(educationRepo *repositories.EducationRepository) *EducationService {
	return &EducationService{
		educationRepository: educationRepo,
	}
}

func (s *EducationService) CreateEducation(userId primitive.ObjectID, education models.Education) error {
	education.UserId = userId
	return s.educationRepository.CreateEducation(education)
}

func (s *EducationService) GetUserEducation(userID primitive.ObjectID) ([]models.Education, error) {
	return s.educationRepository.GetEducationByUserID(userID)
}

func (s *EducationService) UpdateEducation(userId primitive.ObjectID, education models.Education) error {
	if education.UserId != userId {
		return errors.New("unauthorized")
	}

	return s.educationRepository.UpdateEducation(education)
}

func (s *EducationService) DeleteEducatationById(userId primitive.ObjectID, educationId primitive.ObjectID) error {
	// Fetch the education by educationId, not userId
	education, err := s.educationRepository.GetEducationById(educationId)
	if err != nil {
		return errors.New("education not found")
	}

	// Check if the user is authorized to delete this education
	if education.UserId != userId {
		return errors.New("unauthorized")
	}

	// Proceed to delete the education record
	return s.educationRepository.DeleteEducation(educationId)

}
