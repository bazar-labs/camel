package service

import (
	"context"

	"github.com/el-goblino-foundation/turron/internal/blockchain"
	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/ethereum/go-ethereum/common"
)

type Service struct {
	cfg        *config.Service
	store      IStore
	blockchain IBlockchain
}

type IStore interface {
	// Game
	ListGames(ctx context.Context, userID int) ([]domain.Game, error)
	GetGame(ctx context.Context, userID int, gameID int) (*domain.Game, error)
	CreateGame(ctx context.Context, userID int, name string) (*domain.Game, error)

	// Game Contract Addresses
	GetGameContractAddresses(ctx context.Context, gameID int) (*domain.GameContractAddresses, error)
	UpdateGameContractAddresses(ctx context.Context, gameID int, addresses domain.GameContractAddresses) error
}

type IBlockchain interface {
	Economy() blockchain.IEconomy
	InventoryRegistry(address common.Address) blockchain.IInventoryRegistry
}

func New(cfg *config.Service, store IStore, blockchain IBlockchain) *Service {
	return &Service{cfg, store, blockchain}
}
