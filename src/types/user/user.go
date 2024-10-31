﻿package user

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"strings"
)

type User struct {
	ID       uuid.UUID `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
}

func NewUser(name, email, password string) User {
	return User{uuid.New(), name, email, password}
}

type Register struct {
	UserName        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

func (u Register) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.By(func(value interface{}) error {
			if strings.Compare(u.Password, u.ConfirmPassword) != 0 {
				return errors.New("password and confirm password should match")
			}
			return nil
		})),
	)
}

type ChangePassword struct {
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

func (u ChangePassword) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Password, validation.Required, validation.By(func(value interface{}) error {
			if strings.Compare(u.Password, u.ConfirmPassword) != 0 {
				return errors.New("")
			}
			return nil
		})),
	)
}