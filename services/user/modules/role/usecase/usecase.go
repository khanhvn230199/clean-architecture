package usecase

import (
	"context"
	"github.com/example-golang-projects/clean-architecture/services/user/entities"
	"github.com/example-golang-projects/clean-architecture/services/user/modules/role/repository"
	"github.com/jackc/pgx/v5"
)

type RoleRepoForRoleUseCase interface {
	CreateRole(context.Context, entities.Role) error
}

type RoleUseCase struct {
	db          *pgx.Conn
	roleUseCase RoleRepoForRoleUseCase
}

func NewRoleUseCase(db *pgx.Conn) RoleUseCase {
	return RoleUseCase{
		db:          db,
		roleUseCase: repository.NewRoleRepository(db),
	}
}

func (u *RoleUseCase) CreateRole(ctx context.Context, args int) (err error) {
	return u.roleUseCase.CreateRole(ctx, entities.Role{})
}
