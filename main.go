package main

import (
	"os"

	"github.com/el-goblino-foundation/turron/api"
	"github.com/el-goblino-foundation/turron/config"
	"github.com/el-goblino-foundation/turron/db"
	"github.com/el-goblino-foundation/turron/service"
	"github.com/el-goblino-foundation/turron/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load()

	if cfg.App.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Level(zerolog.DebugLevel)
		log.Debug().Interface("config", cfg).Msg("loaded config")
	}

	var (
		db      = db.New(&cfg.Database)
		store   = store.New(db)
		service = service.New(store)
		api     = api.New(&cfg.API, service)
	)

	log.Fatal().Err(api.ListenAndServe()).Msg("server crashed")
}
