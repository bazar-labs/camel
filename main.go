package main

import (
	"fmt"
	"os"

	"github.com/el-goblino-foundation/turron/config"
	"github.com/el-goblino-foundation/turron/db"
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

	db := db.New(&cfg.Database)
	_ = db

	fmt.Println("hi^^")
}
