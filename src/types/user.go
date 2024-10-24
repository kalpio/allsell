package types

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"strings"
)

type UserRegister struct {
	UserName        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

func (u UserRegister) Validate() error {
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
