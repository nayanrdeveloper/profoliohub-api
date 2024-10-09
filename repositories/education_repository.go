package repositories

import (
	"context"
	"profoliohub-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EducationRepository struct {
	collection *mongo.Collection
}

func NewEducationRepository(db *mongo.Database) *EducationRepository {
	return &EducationRepository{
		collection: db.Collection("education"),
	}
}

func (r *EducationRepository) CreateEducation(education models.Education) error {
	_, err := r.collection.InsertOne(context.Background(), education)
	return err
}

func (r *EducationRepository) GetEducationByUserID(userID primitive.ObjectID) ([]models.Education, error) {
	var educations []models.Education
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &educations)
	return educations, err
}

func (r *EducationRepository) UpdateEducation(education models.Education) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": education.ID}, bson.M{"$set": education})
	return err
}

func (r *EducationRepository) DeleteEducation(educationID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": educationID})
	return err
}

func (r *EducationRepository) GetEducationById(educationID primitive.ObjectID) (models.Education, error) {
	var education models.Education
	err := r.collection.FindOne(context.Background(), bson.M{"_id": educationID}).Decode(&education)
	if err != nil {
		return education, err
	}
	return education, nil
}
