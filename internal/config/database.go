package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(databaseName string) (*sql.DB, error) {
	// Create a new database connection
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	fmt.Println("Database connection successful")
	return db, nil
}
