package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Education struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Institution string             `bson:"institution" json:"institution"`
	Degree      string             `bson:"degree" json:"degree"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Startyear   int                `bson:"start_year" json:"start_year"`
	Endyear     int                `bson:"end_year" json:"end_year"`
}
