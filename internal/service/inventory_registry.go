package service

import (
	"context"
	"fmt"
	"math/big"
	"mime/multipart"

	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/google/uuid"
)

func (s *Service) GetItemDefinition(ctx context.Context, gameID uuid.UUID, itemDefID *big.Int) (*domain.ItemDefinition, error) {
	item := &domain.ItemDefinition{ID: itemDefID}

	addresses, err := s.store.GetGameContractAddresses(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game contract addresses: %v", err)
	}

	item.URI, err = s.blockchain.InventoryRegistry(addresses.InventoryRegistry).GetItemDefinition(item.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get item definition: %v", err)
	}

	return item, nil
}

func (s *Service) CreateItemDefinition(ctx context.Context, gameID uuid.UUID, form *multipart.Form) (*domain.ItemDefinition, error) {
	var err error
	itemDef := &domain.ItemDefinition{}

	// TODO better name than hash? also is this just cid or full ipfs url?
	itemDef.URI, err = s.UploadToIPFS(ctx, form)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	addresses, err := s.store.GetGameContractAddresses(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game contract addresses: %v", err)
	}

	itemDef.ID, err = s.blockchain.InventoryRegistry(addresses.InventoryRegistry).CreateItemDefinition(itemDef.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to create item definition: %v", err)
	}

	return itemDef, nil
}
