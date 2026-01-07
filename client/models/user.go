package models

import (
	"github.com/adk-saugat/stash/client/utils"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(username, email, password string) *User {
	return &User{
		ID:       utils.GenerateUUID(),
		Username: username,
		Email:    email,
		Password: password,
	}
}