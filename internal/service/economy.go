package service

import (
	"context"
	"fmt"

	"github.com/bazar-labs/camel/internal/domain"
	"github.com/rs/zerolog/log"
)

func (s *Service) ListGameEconomy(ctx context.Context, userID, gameID int64) ([]domain.GameEconomy, error) {
	return s.store.ListGameEconomy(ctx, userID, gameID)
}

func (s *Service) GetGameEconomy(ctx context.Context, userID, gameID, economyID int64) (*domain.GameEconomy, error) {
	return s.store.GetGameEconomy(ctx, userID, gameID, economyID)
}

func (s *Service) CreateGameEconomy(ctx context.Context, userID, gameID int64, chainNetwork domain.ChainNetwork) (*domain.GameEconomy, error) {
	mastercontracts, err := s.store.GetMasterEconomyContracts(ctx, chainNetwork)
	if err != nil {
		return nil, fmt.Errorf("failed to get master contracts: %v", err)
	}

	gamecontracts, err := s.blockchain.Economy().Deploy(*mastercontracts)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy game economy: %v", err)
	}

	economy, err := s.store.CreateGameEconomy(ctx, userID, gameID, chainNetwork, gamecontracts)
	if err != nil {
		return nil, fmt.Errorf("failed to set game contract addresses: %v", err)
	}

	log.Debug().
		Int64("id", economy.ID).
		Interface("contracts", economy.Contracts).
		Msg("deployed game economy")

	return economy, nil
}
