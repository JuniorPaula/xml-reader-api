package main

import (
	"log"
	"xml-reader-api/internal/config"
)

func main() {
	// Create a new database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Automatically migrate the database
	config.AutoMigrate(db)
}
