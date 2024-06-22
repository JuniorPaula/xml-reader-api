package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// AutoMigrate is a function that automatically migrates the database
// It is used to create the tables in the database
// It is called in the main.go file
func AutoMigrate(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create the users table
	_, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated successfully")
}
