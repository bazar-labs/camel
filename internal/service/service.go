package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/el-goblino-foundation/turron/internal/transact"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Service struct {
	store    IStore
	transact ITransact
}

type IStore interface {
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
}

type ITransact interface {
	Counter(address common.Address) transact.ICounter
}

func New(store IStore, transact ITransact) *Service {
	return &Service{store, transact}
}

func (s *Service) ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	return s.store.ListGames(ctx, userID)
}

// FIXME
var ContractAddress = common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

func (s *Service) GetNumber() (*big.Int, error) {
	number, err := s.transact.Counter(ContractAddress).GetNumber()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the 'number' variable: %v", err)
	}

	return number, nil
}

func (s *Service) SetNumber(number int) error {
	_, err := s.transact.Counter(ContractAddress).SetNumber(number)
	if err != nil {
		return fmt.Errorf("failed to increment the number: %v", err)
	}

	return nil
}
