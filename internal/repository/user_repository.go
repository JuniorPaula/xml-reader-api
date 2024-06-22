package repository

import (
	"context"
	"database/sql"
	"time"
	"xml-reader-api/internal/entity"
)

type UserInterface interface {
	CreateUser(name, email, password string) (int64, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(name, email, password string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.DB.ExecContext(ctx, `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *User) GetUserByEmail(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := u.DB.QueryRowContext(ctx, `SELECT id, name, email, password FROM users WHERE email = ?`, email)
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
