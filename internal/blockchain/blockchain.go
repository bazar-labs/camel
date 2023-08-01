package blockchain

import (
	"crypto/ecdsa"

	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type Client struct {
	client                  *ethclient.Client
	pk                      *ecdsa.PrivateKey
	masterContractAddresses *config.MasterContractAddresses
}

func New(cfg *config.Blockchain) *Client {
	client, err := ethclient.Dial(cfg.URL)
	if err != nil {
		log.Fatal().Msgf("failed to connect to with the client client: %v", err)
	}

	pk, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal().Msgf("failed to parse private key: %v", err)
	}

	return &Client{client, pk, &cfg.MasterContractAddresses}
}
