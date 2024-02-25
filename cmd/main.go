package main

import (

	// "github.com/gin-gonic/gin"

	"github.com/SandeepTukkunor/loan_management_system/routes"
	_ "github.com/lib/pq"
)

func main() {

	r := routes.SetupRouter()

	// Start the server
	r.Run()

}
