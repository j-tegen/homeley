package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	OwnerRole    UserRole = "owner"
	ResidentRole UserRole = "resident"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	User      CreateUserPayload        `json:"user"`
	Household RegisterHouseholdPayload `json:"household"`
}

type CreateUserPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	HouseholdID primitive.ObjectID `json:"householdId"`
	Role        UserRole           `json:"role"`
}
