package models

import "fmt"

type User struct{
	ID 			int64 	`json:"id"`
	Email 		string 	`json:"email"`
	Password 	string	`json:"password`	
}

func NewUser(id int64, email string, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

func (user *User) LoginUser(){
	fmt.Println("user logged in")
}

func (user *User) RegisterUser(){
	fmt.Println("user registered")
}