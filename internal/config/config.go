package config

import (
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	App      App
	API      API
	Database Database
}

type App struct {
	Environment string `env:"APP_ENVIRONMENT" env-required:"true"`
}

type API struct {
	ServerPort int `env:"API_SERVER_PORT" env-required:"true"`
	Pinata     Pinata
}

type Pinata struct {
	JWT string `env:"API_PINATA_JWT" env-required:"true"`
}

type Database struct {
	URL string `env:"DATABASE_URL" env-required:"true"`
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
