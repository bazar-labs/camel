package service

import (
	"context"
	"fmt"
	"math/big"

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
}

type IBlockchain interface {
	Counter(address common.Address) blockchain.ICounter
}

func New(store IStore, blockchain IBlockchain) *Service {
	return &Service{store, blockchain}
}

func (s *Service) ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	return s.store.ListGames(ctx, userID)
}

// FIXME
var ContractAddress = common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

func (s *Service) GetNumber() (*big.Int, error) {
	number, err := s.blockchain.Counter(ContractAddress).GetNumber()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the 'number' variable: %v", err)
	}

	return number, nil
}

func (s *Service) SetNumber(number int) error {
	_, err := s.blockchain.Counter(ContractAddress).SetNumber(number)
	if err != nil {
		return fmt.Errorf("failed to increment the number: %v", err)
	}

	return nil
}
