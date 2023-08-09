package main

import (
	"fmt"
	"os"

	"github.com/bazar-labs/turron/internal/blockchain"
	"github.com/bazar-labs/turron/internal/config"
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

	blockchain := blockchain.New(&cfg.Blockchain)

	addresses, err := blockchain.Economy().Deploy()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to deploy economy")
	}
	fmt.Printf("Economy deployed at: %#v\n", addresses)

	itemDefID, err := blockchain.InventoryRegistry(addresses.InventoryRegistry).CreateItem("https://example.com/item-def-1")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create item definition")
	}
	fmt.Printf("Item Definition ID: %s\n", itemDefID)

	itemDefURI, err := blockchain.InventoryRegistry(addresses.InventoryRegistry).GetItem(itemDefID)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get item definition")
	}
	fmt.Printf("Item Definition URI: %s\n", itemDefURI)
}
