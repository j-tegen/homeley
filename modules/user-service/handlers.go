package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/j-tegen/homeley/shared/models"
	"github.com/j-tegen/homeley/shared/db"
)


func UserRegisterHandler(c *gin.Context, connection *db.DbConnection) {
	var registerRequest models.RegisterPayload 
	if err := c.BindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := ValidateRegisterPayload(registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	user, err := Register(registerRequest, connection.Client)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, user)
}

func UserLoginHandler(c *gin.Context) {
	var loginRequest models.LoginPayload
	if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	token, err := CallAuthLogin(loginRequest)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func MeHandler(c *gin.Context, connection *db.DbConnection) {
	userID, exists := c.Get("UserId")
	if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
	}
	user, err := GetUser(userID.(string), connection.Client)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, user)
}
