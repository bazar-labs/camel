package config

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	App        App
	API        API
	Database   Database
	Service    Service
	Blockchain Blockchain
}

type App struct {
	Environment string `env:"APP_ENVIRONMENT" env-required:"true"`
}

type API struct {
	ServerPort int `env:"API_SERVER_PORT" env-required:"true"`
}

type Database struct {
	DSN string `env:"DATABASE_DSN" env-required:"true"`
}

type Service struct {
	Pinata Pinata
}

type Pinata struct {
	JWT string `env:"SERVICE_PINATA_JWT" env-required:"true"`
}

type Blockchain struct {
	URL                     string `env:"BLOCKCHAIN_URL" env-required:"true"`
	PrivateKey              string `env:"BLOCKCHAIN_PRIVATE_KEY" env-required:"true"`
	MasterContractAddresses MasterContractAddresses
}

type MasterContractAddresses struct {
	BoringFactory     common.Address `env:"BLOCKCHAIN_MASTER_CONTRACT_ADDRESSES_BORING_FACTORY" env-required:"true"`
	InventoryRegistry common.Address `env:"BLOCKCHAIN_MASTER_CONTRACT_ADDRESSES_INVENTORY_REGISTRY" env-required:"true"`
}

func Load(paths ...string) *Config {
	err := godotenv.Load(paths...)
	if err != nil {
		log.Warn().Msgf("failed to load env vars from %s file(s): %v", strings.Join(paths, ","), err)
	}

	var cfg Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal().Msgf("failed to load config: %v", err)
	}

	return &cfg
}
