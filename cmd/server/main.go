package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/el-goblino-foundation/turron/contract"
	"github.com/el-goblino-foundation/turron/internal/api"
	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/el-goblino-foundation/turron/internal/db"
	"github.com/el-goblino-foundation/turron/internal/service"
	"github.com/el-goblino-foundation/turron/internal/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	_ = api

	testContract()

	// log.Fatal().Err(api.ListenAndServe()).Msg("server crashed")
}

var PrivateKey, _ = crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
var ContractAddress = common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

// constants:
// - private key
// - chain id
// - rpc endpoint
// - xyz contract address

func testContract() {
	// Create an IPC based RPC connection t`o a remote node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal().Msgf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	instance, err := contract.NewCounterContract(ContractAddress, client)
	if err != nil {
		log.Fatal().Msgf("Failed to instantiate a Token contract: %v", err)
	}

	fmt.Printf("instance %#v\n", instance)

	// Call the 'number' public variable getter
	number, err := instance.Number(&bind.CallOpts{})
	if err != nil {
		log.Fatal().Msgf("Failed to retrieve the 'number' variable: %v", err)
	}
	fmt.Println("Current Number:", number)

	// Call the 'increment' method
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(PrivateKey, new(big.Int).SetInt64(31337))
	tx, err := instance.Increment(transactOpts)
	if err != nil {
		log.Fatal().Msgf("Failed to increment the number: %v", err)
	}
	fmt.Printf("Incremented, transaction hash: %s\n", tx.Hash().Hex())

	// Call the 'setNumber' method
	newNumber := big.NewInt(10)
	tx, err = instance.SetNumber(transactOpts, newNumber)
	if err != nil {
		log.Fatal().Msgf("Failed to set a new number: %v", err)
	}
	fmt.Printf("Set new number, transaction hash: %s\n", tx.Hash().Hex())
}
