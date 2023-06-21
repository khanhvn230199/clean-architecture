package dependency

import (
	"context"
	"fmt"
	"github.com/example-golang-projects/clean-architecture/cmd/server/user/config"
	"github.com/example-golang-projects/clean-architecture/packages/database"
	"github.com/example-golang-projects/clean-architecture/packages/database/migration"
	"github.com/jackc/pgx/v5"
	"log"
)

type UserDependency struct {
	Config config.Config
	DB     *pgx.Conn

	// List of client/third-party
}

func InitUserDependency(cfg config.Config) UserDependency {
	connStr := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", cfg.Database.UserDB.Database, cfg.Database.UserDB.Username, cfg.Database.UserDB.Password, cfg.Database.UserDB.Host, cfg.Database.UserDB.Port, cfg.Database.UserDB.SSLMode)
	ctx := context.Background()
	db, err := database.NewDatabase(ctx, connStr)
	if err != nil {
		log.Panic(err)
	}

	err = migration.Up(ctx, db, "./services/user/migrations", cfg.ServiceName)
	if err != nil {
		log.Panic(err)
	}

	return UserDependency{
		Config: cfg,
		DB:     db,
	}
}
