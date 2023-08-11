package service

import (
	"context"

	"github.com/bazar-labs/turron/internal/blockchain"
	"github.com/bazar-labs/turron/internal/config"
	"github.com/bazar-labs/turron/internal/domain"
	"github.com/ethereum/go-ethereum/common"
)

type Service struct {
	cfg        *config.Service
	store      IStore
	blockchain IBlockchain
}

type IStore interface {
	// Game
	ListGame(ctx context.Context, userID int64) ([]domain.Game, error)
	GetGame(ctx context.Context, userID int64, gameID int64) (*domain.Game, error)
	CreateGame(ctx context.Context, userID int64, name string) (*domain.Game, error)

	// Game Economy
	ListGameEconomy(ctx context.Context, userID, gameID int64) ([]domain.GameEconomy, error)
	GetGameEconomy(ctx context.Context, userID, gameID, economyID int64) (*domain.GameEconomy, error)
	CreateGameEconomy(ctx context.Context, userID, gameID int64, chainNetwork domain.ChainNetwork, addresses domain.GameEconomyContractAddresses) (*domain.GameEconomy, error)

	// Master Contracts
	GetMasterContracts(ctx context.Context, chainNetwork domain.ChainNetwork) (*domain.MasterContracts, error)
}

type IBlockchain interface {
	Economy() blockchain.IEconomy
	InventoryRegistry(address common.Address) blockchain.IInventoryRegistry
}

func New(cfg *config.Service, store IStore, blockchain IBlockchain) *Service {
	return &Service{cfg, store, blockchain}
}
