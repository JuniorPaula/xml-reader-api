package main

import (
	"fmt"
	"log"
	"net/http"
	"xml-reader-api/internal/config"
	"xml-reader-api/internal/routes"
)

func main() {
	cfg, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal(err)
	}
	// Create a new database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Automatically migrate the database
	config.AutoMigrate(db)

	routes := routes.NewRoutes(db, cfg)

	// Listen and serve the API
	fmt.Printf("Server running on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, routes))
}
