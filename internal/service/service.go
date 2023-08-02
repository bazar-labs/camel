package service

import (
	"context"

	"github.com/el-goblino-foundation/turron/internal/blockchain"
	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Service struct {
	store      IStore
	blockchain IBlockchain
}

type IStore interface {
	// Game
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
	GetGame(ctx context.Context, userID uuid.UUID, gameID uuid.UUID) (*domain.Game, error)
	CreateGame(ctx context.Context, userID uuid.UUID, name string) (*domain.Game, error)

	// Game Contract Addresses
	GetGameContractAddresses(ctx context.Context, gameID uuid.UUID) (*domain.GameContractAddresses, error)
	UpdateGameContractAddresses(ctx context.Context, gameID uuid.UUID, addresses domain.GameContractAddresses) error
}

type IBlockchain interface {
	Economy() blockchain.IEconomy
	InventoryRegistry(address common.Address) blockchain.IInventoryRegistry
}

func New(store IStore, blockchain IBlockchain) *Service {
	return &Service{store, blockchain}
}
