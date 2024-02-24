package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/SandeepTukkunor/loan_management_system/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database", conn)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
