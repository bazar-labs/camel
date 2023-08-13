package service

import (
	"context"
	"fmt"
	"math/big"
	"mime/multipart"

	"github.com/bazar-labs/turron/internal/domain"
)

func (s *Service) ListItem(ctx context.Context, userID, gameID, economyID int64) ([]domain.Item, error) {
	// TODO implement this
	panic("'ListItem' not implemented")
}

func (s *Service) GetItem(ctx context.Context, userID, gameID, economyID int64, itemDefID *big.Int) (*domain.Item, error) {
	item := &domain.Item{DefID: itemDefID}

	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game economy: %v", err)
	}

	// TODO not found error
	item.URI, err = s.blockchain.InventoryRegistry(economy.Contracts.InventoryRegistry.Address).GetItem(item.DefID)
	if err != nil {
		return nil, fmt.Errorf("failed to get item definition: %v", err)
	}

	return item, nil
}

func (s *Service) CreateItem(ctx context.Context, userID, gameID, economyID int64, form *multipart.Form) (*domain.Item, error) {
	var err error
	itemdef := &domain.Item{}

	// TODO better name than hash? also is this just cid or full ipfs url?
	itemdef.URI, err = s.UploadToIPFS(ctx, form)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game economy: %v", err)
	}

	itemdef.DefID, err = s.blockchain.InventoryRegistry(economy.Contracts.InventoryRegistry.Address).CreateItem(itemdef.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to create item definition: %v", err)
	}

	return itemdef, nil
}
