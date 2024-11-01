package main

import (
    "context"
    "errors"
    "github.com/j-tegen/homeley/shared/models"
    "github.com/j-tegen/homeley/shared/utils" 
    "fmt"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

func AuthenticateUser(payload models.LoginPayload, client *mongo.Client) (models.User, error) {
    collection := client.Database("homeley").Collection("users")
    var user models.User

    err := collection.FindOne(context.TODO(), bson.M{"email": payload.Email}).Decode(&user)
    fmt.Println("AuthenticateUser", user, payload.Password, user.Password)
    if err != nil {
        return  models.User{}, err
    }

    if !utils.ValidatePassword(user.Password, payload.Password) {
        fmt.Println("Invalid credentials")
        return models.User{}, errors.New("Invalid credentials")
    }

    return user, nil
}
