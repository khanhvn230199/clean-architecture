package dependency

import (
	"context"
	"fmt"
	"github.com/example-golang-projects/clean-architecture/cmd/server/master/config"
	"github.com/example-golang-projects/clean-architecture/packages/database"
	"github.com/example-golang-projects/clean-architecture/packages/database/migration"
	"github.com/jackc/pgx/v5"
	"log"
)

type MasterDependency struct {
	Config config.Config
	DB     *pgx.Conn

	// List of client/third-party
}

func InitMasterDependency(cfg config.Config) MasterDependency {
	connStr := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", cfg.Database.MasterDB.Database, cfg.Database.MasterDB.Username, cfg.Database.MasterDB.Password, cfg.Database.MasterDB.Host, cfg.Database.MasterDB.Port, cfg.Database.MasterDB.SSLMode)
	ctx := context.Background()
	db, err := database.NewDatabase(ctx, connStr)
	if err != nil {
		log.Panic(err)
	}

	err = migration.Up(ctx, db, "./services/master/migrations", cfg.ServiceName)
	if err != nil {
		log.Panic(err)
	}
	return MasterDependency{
		Config: cfg,
		DB:     db,
	}
}
