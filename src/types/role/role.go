package role

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

var (
	Administrator = NewRole("Administrator")
	User          = NewRole("User")
)

type Role struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

func NewRole(name string) *Role {
	return &Role{
		ID:   uuid.New(),
		Name: name,
	}
}

func (r Role) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(2, 50)))
}

type UserRole struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	RoleID uuid.UUID `json:"role_id" db:"role_id"`
}

func NewUserRole(userID, roleID uuid.UUID) *UserRole {
	return &UserRole{
		UserID: userID,
		RoleID: roleID,
	}
}
