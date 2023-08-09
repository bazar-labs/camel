package service

import (
	"context"
	"fmt"

	"github.com/bazar-labs/turron/internal/domain"
	"github.com/rs/zerolog/log"
)

func (s *Service) ListGameEconomy(ctx context.Context, userID, gameID int64) ([]domain.GameEconomy, error) {
	return s.store.ListGameEconomy(ctx, userID, gameID)
}

func (s *Service) GetGameEconomy(ctx context.Context, userID, gameID, economyID int64) (*domain.GameEconomy, error) {
	return s.store.GetGameEconomy(ctx, userID, gameID, economyID)
}

func (s *Service) CreateGameEconomy(ctx context.Context, userID, gameID int64, chainNetwork domain.ChainNetwork) (*domain.GameEconomy, error) {
	// TODO add check not to deploy twice

	addresses, err := s.blockchain.Economy().Deploy()
	if err != nil {
		return nil, fmt.Errorf("failed to deploy game economy: %v", err)
	}

	economy, err := s.store.CreateGameEconomy(ctx, userID, gameID, chainNetwork, addresses)
	if err != nil {
		return nil, fmt.Errorf("failed to set game contract addresses: %v", err)
	}

	log.Debug().
		Int64("game_economy_id", economy.ID).
		Interface("game_contract_addresses", economy.ContractAddresses).
		Msg("deployed game economy on chain")

	return economy, nil
}
