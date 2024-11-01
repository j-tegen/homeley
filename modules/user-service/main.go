package main

import (
    "github.com/gin-gonic/gin"
    "os"

    "github.com/j-tegen/homeley/shared/db"
    "github.com/j-tegen/homeley/shared/middleware"
)


func main() {
    connection := db.Connect()
    router := gin.Default()

    public := router.Group("/")
    public.POST("/users/login", func(c *gin.Context) {
        UserLoginHandler(c)
    })
    public.POST("/users/register", func(c *gin.Context) {
        UserRegisterHandler(c, connection)
    })

    protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/users/me", func(c *gin.Context) {
		MeHandler(c, connection)
	})

    router.Run("0.0.0.0:" + os.Getenv("PORT"))

}
