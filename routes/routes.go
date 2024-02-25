package routes

import (
	"github.com/SandeepTukkunor/loan_management_system/middleware"
	"github.com/SandeepTukkunor/loan_management_system/views/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// User routes
	r.POST("/signup", user.SignUp)
	r.POST("/login", user.Login)
	r.GET("/validate", middleware.RequireAuth, user.ValidateToken)
	r.POST("/logout", user.Logout)

	return r
}
