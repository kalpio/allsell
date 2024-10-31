package user

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/kalpio/allsell/src/types/time"
	"strings"
)

type User struct {
	ID                 uuid.UUID    `json:"id" db:"id"`
	Name               string       `json:"name" db:"name"`
	Email              string       `json:"email" db:"email"`
	Password           string       `json:"password" db:"password"`
	LastPasswordChange *time.DbTime `json:"last_password_change" db:"last_password_change"`
}

func NewUser(name, email, password string) User {
	return User{uuid.New(), name, email, password, time.Now()}
}

type Register struct {
	UserName        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm-password"`
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
	CurrentPassword    string `form:"current-password"`
	NewPassword        string `form:"new-password"`
	ConfirmNewPassword string `form:"confirm-new-password"`
}

func (u ChangePassword) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.CurrentPassword, validation.Required),
		validation.Field(&u.NewPassword, validation.By(func(value interface{}) error {
			if strings.Compare(u.NewPassword, u.ConfirmNewPassword) != 0 {
				return errors.New("password and confirm password should match")
			}
			return nil
		})),
	)
}
