package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "fmt"

    "github.com/j-tegen/homeley/shared/db"
    "github.com/j-tegen/homeley/shared/models"
    "github.com/j-tegen/homeley/shared/auth"
)

func main() {
    connection := db.Connect()

    r := gin.Default() 
    r.POST("/jwt", func(context *gin.Context) {
        jwtHandler(context, connection)
    })
    r.Run("0.0.0.0:" + os.Getenv("PORT"))
}

func jwtHandler(c *gin.Context, connection *db.DbConnection) {
    var loginRequest models.LoginPayload
    if err := c.BindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("jwtHandler", loginRequest)
    user, err := AuthenticateUser(loginRequest, connection.Client)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }

    token, err := auth.GenerateJWT(user.ID.Hex(), user.HouseholdID.Hex())
    if err != nil { 

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})    
}
