package repository

import (
	"context"
	"database/sql"
	"time"
)

type UserInterface interface {
	CreateUser(name, email, password string) error
}

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.DB.ExecContext(ctx, `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		return err
	}
	return nil
}
