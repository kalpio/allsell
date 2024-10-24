package services

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/models"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{db}
}

func (u UserService) Register(ctx context.Context, user models.User) error {
	sql := `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?);`
	if _, err := u.db.ExecContext(ctx, sql, user.ID.String(), user.Name, user.Email, user.Password); err != nil {
		return fmt.Errorf("user:register: %w", err)
	}

	return nil
}
