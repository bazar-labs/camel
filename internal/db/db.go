package db

import (
	"database/sql"

	"github.com/el-goblino-foundation/turron/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func New(cfg *config.Database) *sql.DB {
	log.Debug().Msg("connecting to database...")

	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.Fatal().Msgf("can't connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Msgf("can't connect to database: %v", err)
	}

	return db
}
