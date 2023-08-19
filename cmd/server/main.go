package main

import (
	"os"

	"github.com/bazar-labs/camel/internal/api"
	"github.com/bazar-labs/camel/internal/blockchain"
	"github.com/bazar-labs/camel/internal/config"
	"github.com/bazar-labs/camel/internal/db"
	"github.com/bazar-labs/camel/internal/service"
	"github.com/bazar-labs/camel/internal/store"
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
		db         = db.New(&cfg.Database)
		store      = store.New(db)
		blockchain = blockchain.New(&cfg.Blockchain)
		service    = service.New(&cfg.Service, store, blockchain)
		api        = api.New(&cfg.API, service)
	)

	log.Fatal().Err(api.ListenAndServe()).Msg("server crashed")
}
