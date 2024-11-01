package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func main() {
    r := gin.Default() 
    r.POST("/login", loginHandler)
    r.Run("0.0.0.0:" + os.Getenv("PORT"))
}

func loginHandler(c *gin.Context) {
    var loginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if AuthenticateUser(loginRequest.Username, loginRequest.Password) {
        c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}
