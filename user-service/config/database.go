package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/database/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PsqlDB.User, cfg.PsqlDB.Password, cfg.PsqlDB.Host, cfg.PsqlDB.Port, cfg.PsqlDB.DBName)
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database to host " + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database connection")
		return nil, err
	}

	seeds.SeedRole(db)
	seeds.SeedAdmin(db)

	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpenConns)

	return &Postgres{DB: db}, nil
}
