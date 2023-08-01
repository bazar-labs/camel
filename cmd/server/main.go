package main

import (
	"fmt"
	"os"

	"github.com/el-goblino-foundation/turron/internal/api"
	"github.com/el-goblino-foundation/turron/internal/blockchain"
	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/el-goblino-foundation/turron/internal/db"
	"github.com/el-goblino-foundation/turron/internal/service"
	"github.com/el-goblino-foundation/turron/internal/store"
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
		service    = service.New(store, blockchain)
		api        = api.New(&cfg.API, service)
	)
	_ = api

	addresses, err := blockchain.Economy().Deploy()
	fmt.Printf("addresses %#v\n", addresses)
	fmt.Printf("err %#v\n", err)

}
