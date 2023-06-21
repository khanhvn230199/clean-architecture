package database

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(ctx context.Context, connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
