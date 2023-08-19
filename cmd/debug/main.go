package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bazar-labs/turron/internal/blockchain"
	"github.com/bazar-labs/turron/internal/config"
	"github.com/bazar-labs/turron/internal/domain"
	"github.com/ethereum/go-ethereum/common"
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

	mc := domain.MasterEconomyContracts{
		BoringFactory:       common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		InventoryRegistry:   common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
		InventoryController: common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		Behaviors: domain.BehaviourContractInfo{
			PurchaseItemWithETH: common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
		},
	}

	data, _ := json.Marshal(mc)
	fmt.Println(string(data))

	addresses, err := blockchain.Economy().Deploy(mc)
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
