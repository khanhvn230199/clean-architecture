package repository

import (
	"context"
	"fmt"
	"github.com/example-golang-projects/clean-architecture/services/user/entities"
	"github.com/jackc/pgx/v5"
)

type RoleRepository struct {
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role entities.Role) error {
	fmt.Println("RoleRepository")
	return nil
}
