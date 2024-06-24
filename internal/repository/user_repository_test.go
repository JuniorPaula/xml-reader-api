package repository

import (
	"database/sql"
	"testing"
	"xml-reader-api/internal/config"
)

func initDatabase() (*sql.DB, error) {
	// Create a new database connection
	db, err := config.ConnectDatabase("./xml_reader_test.db")
	if err != nil {
		return nil, err
	}
	// Automatically migrate the database
	config.AutoMigrate(db)
	return db, nil
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}

func TestUserRepository(t *testing.T) {
	DB, err := initDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer DB.Close()

	userRepo := NewUserRepository(DB)
	id, err := userRepo.CreateUser("James Foo", "james@foo.com", "1234")
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	if id != 1 {
		t.Fatalf("Expected user ID to be 1, got %d", id)
	}

	user, err := userRepo.GetUserByEmail("james@foo.com")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if user.Name != "James Foo" {
		t.Errorf("Expected user name to be James Foo, got %s", user.Name)
	}

	if err := dropTable(DB); err != nil {
		t.Fatalf("Failed to drop table: %v", err)
	}
}
