package models

import (
	"context"

	"github.com/adk-saugat/stash/server/pkg/database"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// FindByEmail retrieves a user by email, returns nil if not found
func (u *User) FindByEmail(email string) error {
	return database.Pool.QueryRow(context.Background(),
		"SELECT id, username, email, password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
}

// Create inserts a new user into the database
func (u *User) Create() error {
	_, err := database.Pool.Exec(context.Background(),
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		u.Username, u.Email, u.Password,
	)
	return err
}