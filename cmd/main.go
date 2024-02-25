package main

import (
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"

	"github.com/SandeepTukkunor/loan_management_system/internal/db"
	"github.com/SandeepTukkunor/loan_management_system/routes"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database", conn)

	r := routes.SetupRouter()

    // Start the server
    r.Run()

}
