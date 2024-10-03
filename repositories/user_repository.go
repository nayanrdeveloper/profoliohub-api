package repositories

import (
	"context"
	"profoliohub-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositories struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepositories {
	return &UserRepositories{
		collection: db.Collection("users"),
	}
}

func (r *UserRepositories) CreateUser(user models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepositories) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}
