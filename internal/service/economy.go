package service

import (
	"context"
	"fmt"

	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/google/uuid"
)

func (s *Service) DeployGameEconomy(ctx context.Context, gameID uuid.UUID) (domain.GameContractAddresses, error) {
	// TODO add check not to deploy twice

	addresses, err := s.blockchain.Economy().Deploy()
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to deploy game economy: %v", err)
	}

	err = s.store.UpdateGameContractAddresses(ctx, gameID, addresses)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to update game contract addresses: %v", err)
	}

	return addresses, nil
}
