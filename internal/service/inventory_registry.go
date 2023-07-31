package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
)

func (s *Service) GetItemDefinition(ctx context.Context, gameID uuid.UUID, itemDefID int) (string, error) {
	addresses, err := s.store.GetGameContractAddresses(ctx, gameID)
	if err != nil {
		return "", fmt.Errorf("failed to get game contract addresses: %v", err)
	}

	itemDef, err := s.blockchain.InventoryRegistry(addresses.InventoryRegistry).GetItemDefinition(big.NewInt(int64(itemDefID)))
	if err != nil {
		return "", fmt.Errorf("failed to get item definition: %v", err)
	}

	return itemDef, nil
}

func (s *Service) CreateItemDefinition(ctx context.Context, gameID uuid.UUID, itemDefURI string) (*types.Transaction, error) {
	addresses, err := s.store.GetGameContractAddresses(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game contract addresses: %v", err)
	}

	tx, err := s.blockchain.InventoryRegistry(addresses.InventoryRegistry).CreateItemDefinition(itemDefURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create item definition: %v", err)
	}

	return tx, nil
}
