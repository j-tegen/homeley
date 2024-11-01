package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/j-tegen/homeley/shared/config"
)

type Claims struct {
	UserID      string `json:"userId"`
	HouseholdID string `json:"householdId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, householdID string) (string, error) {
	claims := &Claims{
		UserID:      userID,
		HouseholdID: householdID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}
