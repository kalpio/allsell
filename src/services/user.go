package services

import (
	"context"
	"fmt"
	"github.com/kalpio/allsell/src/types"

	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/models"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{db}
}

func (u UserService) Register(ctx context.Context, user models.User) error {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	sql := `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?);`
	if _, err := u.db.ExecContext(ctx, sql, user.ID.String(), user.Name, user.Email, passwordHash); err != nil {
		return fmt.Errorf("user:register: failed to insert user to database %w", err)
	}

	return nil
}

func (u UserService) Login(ctx context.Context, username, password string) (types.LoginResult, *models.User) {
	user := models.User{}
	sql := "select * from users where name = ? limit 1;"
	if err := u.db.QueryRowxContext(ctx, sql, username).StructScan(&user); err != nil {
		return types.LoginFailed(types.AuthenticationFailed), nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return types.LoginFailed(types.AuthenticationFailed), nil
	}

	return types.LoginSuccess(), &user
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("user:register: failed to create password hash: %w", err)
	}

	return string(bytes), nil
}
