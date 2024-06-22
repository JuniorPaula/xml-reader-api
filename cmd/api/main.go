package main

import (
	"fmt"
	"log"
	"net/http"
	"xml-reader-api/internal/config"
	"xml-reader-api/internal/routes"
)

const port = "6969"

func main() {
	// Create a new database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Automatically migrate the database
	config.AutoMigrate(db)

	routes := routes.NewRoutes(db)

	// Listen and serve the API
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, routes))
}
