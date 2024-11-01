package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Household struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name"`
}

type RegisterHouseholdPayload struct {
	Name string             `json:"name"`
	ID   primitive.ObjectID `bson:"_id"`
}
