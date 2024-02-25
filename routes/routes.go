package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SandeepTukkunor/loan_management_system/views/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", user.SignUp)
	r.POST("/login", user.Login)

	return r
}
