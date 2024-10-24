package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
}

func NewUser(name, email, password string) User {
	return User{uuid.New(), name, email, password}
}
