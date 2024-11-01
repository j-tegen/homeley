package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/j-tegen/homeley/shared/models"
	"github.com/j-tegen/homeley/shared/utils"
)

type JWTResponse struct {
	Token string `json:"token"`
}

type householdIdChannel struct {
	Value primitive.ObjectID
	Error error
}

func Register(payload models.RegisterPayload, client *mongo.Client) (models.User, error) {
	houseIdChannel := make(chan householdIdChannel)

	hashedPassword, _ := utils.HashPassword(payload.User.Password)
	go handleRegisterHousehold(payload.Household, client, houseIdChannel)

	var role models.UserRole

	if payload.Household.ID == primitive.NilObjectID {
		role = models.OwnerRole
	} else {
		role = models.ResidentRole
	}

	householdID := <-houseIdChannel

	if householdID.Error != nil {
		return models.User{}, householdID.Error
	}

	user := models.User{
		ID:          primitive.NewObjectID(),
		Email:       payload.User.Email,
		Password:    hashedPassword,
		LastName:    payload.User.LastName,
		FirstName:   payload.User.FirstName,
		HouseholdID: householdID.Value,
		Role:        role,
	}

	collection := client.Database("homeley").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func ValidateRegisterPayload(payload models.RegisterPayload) error {
	if payload.User.Email == "" {
		return fmt.Errorf("email is required")
	}
	if payload.User.Password == "" {
		return fmt.Errorf("password is required")
	}
	if payload.User.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if payload.User.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if payload.Household.Name == "" && payload.Household.ID == primitive.NilObjectID {
		return fmt.Errorf("household name or ID is required")
	}
	return nil
}

func handleRegisterHousehold(payload models.RegisterHouseholdPayload, client *mongo.Client, channel chan householdIdChannel) {
	collection := client.Database("homeley").Collection("households")

	if payload.ID != primitive.NilObjectID {
		var household models.Household
		err := collection.FindOne(context.TODO(), bson.M{"_id": payload.ID}).Decode(&household)
		if err != nil {
			channel <- householdIdChannel{primitive.NilObjectID, err}
			return
		}
		channel <- householdIdChannel{household.ID, nil}
		return
	}

	if payload.Name == "" {
		channel <- householdIdChannel{primitive.NilObjectID, fmt.Errorf("household name is required")}
		return
	}

	household := models.Household{
		ID:   primitive.NewObjectID(),
		Name: payload.Name,
	}

	_, err := collection.InsertOne(context.TODO(), household)
	if err != nil {
		channel <- householdIdChannel{primitive.NilObjectID, err}
		return
	}
	channel <- householdIdChannel{household.ID, nil}
	return
}

func CallAuthLogin(payload models.LoginPayload) (string, error) {
	url := os.Getenv("AUTH_SERVICE_URL") + "/jwt"

	jsonPayload, err := json.Marshal(map[string]string{"email": payload.Email, "password": payload.Password})

	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth-service returned status: %d", resp.StatusCode)
	}

	var jwtResp JWTResponse
	if err := json.NewDecoder(resp.Body).Decode(&jwtResp); err != nil {
		return "", err
	}

	return jwtResp.Token, nil
}

func GetUser(userID string, client *mongo.Client) (models.User, error) {
	collection := client.Database("homeley").Collection("users")
	var user models.User

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
