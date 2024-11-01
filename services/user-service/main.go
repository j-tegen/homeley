package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "log"

    "github.com/j-tegen/homeley/shared/models"
    "github.com/j-tegen/homeley/shared/db"
)


func main() {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set")
    }

    // Initialize the Database
    database, err := db.NewDatabase(mongoURI)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer func() {
        if err := database.Client.Disconnect(nil); err != nil {
            log.Fatalf("Error disconnecting from MongoDB: %v", err)
        }
    }()

    r := gin.Default()
    r.POST("/user/login", userLoginHandler)
    r.Run("0.0.0.0:" + os.Getenv("PORT"))
}

func userLoginHandler(c *gin.Context) {
	var loginRequest models.LoginPayload
    if err := c.BindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    response, err := CallAuthLogin(loginRequest.Username, loginRequest.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.String(http.StatusOK, response)
}
