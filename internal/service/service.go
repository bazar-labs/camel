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
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
	CreateGame(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error)
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
