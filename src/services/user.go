package services

import (
	"context"
	"fmt"
	"github.com/kalpio/allsell/src/types/login"
	"github.com/kalpio/allsell/src/types/time"

	"github.com/kalpio/allsell/src/types/user"

	"github.com/jmoiron/sqlx"
	"github.com/kalpio/option"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{db}
}

func (u UserService) Register(ctx context.Context, user user.User) error {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	sql := `insert into users (id, name, email, password, last_password_change) values (?, ?, ?, ?, ?);`
	if _, err := u.db.ExecContext(ctx, sql, user.ID.String(), user.Name, user.Email, passwordHash, time.Now().ToDb()); err != nil {
		return fmt.Errorf("user:register: failed to insert user to database %w", err)
	}

	return nil
}

func (u UserService) Login(ctx context.Context, username, password string) (login.LoginResult, option.Option[user.User]) {
	usr := user.User{}
	sql := `select * from users where name = ? limit 1;`
	if err := u.db.QueryRowxContext(ctx, sql, username).StructScan(&usr); err != nil {
		return login.LoginFailed(login.AuthenticationFailed), option.None[user.User]()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); err != nil {
		return login.LoginFailed(login.AuthenticationFailed), option.None[user.User]()
	}

	return login.LoginSuccess(), option.Some(usr)
}

func (u UserService) ChangePassword(ctx context.Context, username string, currentPassword, newPassword string) error {
	wrappedUser, err := u.Get(ctx, username)
	if err != nil || wrappedUser.IsNone() {
		return err
	}

	usr := wrappedUser.Unwrap()
	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(currentPassword)); err != nil {
		return err
	}

	hash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}
	sql := `update users set password = ?, last_password_change = ? where name = ?;`
	if _, err := u.db.ExecContext(ctx, sql, hash, time.Now().ToDb(), usr.Name); err != nil {
		return err
	}

	return nil
}

func (u UserService) Get(ctx context.Context, username string) (option.Option[user.User], error) {
	usr := user.User{}
	sql := `select * from users where name = ? limit 1;`
	if err := u.db.QueryRowxContext(ctx, sql, username).StructScan(&usr); err != nil {
		return option.None[user.User](), err
	}

	return option.Some(usr), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("user:register: failed to create password hash: %w", err)
	}

	return string(bytes), nil
}
