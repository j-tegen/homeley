package config

import (
	"os"
	"time"
)

var (
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTExpiry = time.Hour
)
