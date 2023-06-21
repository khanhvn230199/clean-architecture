package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Role struct {
	ID        pgtype.Text
	Name      pgtype.Text
	CreatedAt pgtype.Time
	UpdatedAt pgtype.Time
	DeletedAt pgtype.Time
}
