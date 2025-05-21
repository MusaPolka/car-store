package main

import (
	"log"

	"ecommerce/api-gateway/routes" // adjust the import path as needed
)

func main() {
	// Create a Gin router using the SetupRouter function.
	router := routes.SetupRouter()

	log.Println("API Gateway running at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
