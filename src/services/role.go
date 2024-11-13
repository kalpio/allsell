package services

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/types/role"
	"github.com/kalpio/option"
)

type RoleService struct {
	db *sqlx.DB
}

func NewRoleService(db *sqlx.DB) RoleService {
	return RoleService{db: db}
}

func (srv RoleService) Create(ctx context.Context, name string) option.Option[role.Role] {
	newRole := role.NewRole(name)
	query := `insert into roles (id, name) values (?, ?);`
	_, err := srv.db.ExecContext(ctx, query, newRole.ID, newRole.Name)
	if err != nil {
		return option.None[role.Role](err)
	}

	return option.Some(*newRole)
}
