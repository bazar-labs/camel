package service

import (
	"context"
	"fmt"
)

func (s *Service) GetBehaviorState(ctx context.Context, userID, gameID, economyID int64, behavior string) (bool, error) {
	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return false, err
	}

	switch behavior {
	case "purchase_item":
		return s.blockchain.InventoryController(economy.Contracts.InventoryController.Address).IsInventoryBehavior(economy.Contracts.Behaviors.PurchaseItem.Address)
	default:
		return false, fmt.Errorf("unknown behavior: %s", behavior)
	}
}

func (s *Service) EnableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior string) error {
	return s.setBehaviourState(ctx, userID, gameID, economyID, behavior, true)
}

func (s *Service) DisableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior string) error {
	return s.setBehaviourState(ctx, userID, gameID, economyID, behavior, false)
}

func (s *Service) setBehaviourState(ctx context.Context, userID, gameID, economyID int64, behavior string, state bool) error {
	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return err
	}

	switch behavior {
	case "purchase_item":
		economy.Contracts.Behaviors.PurchaseItem.Enabled = state
		err := s.blockchain.InventoryController(economy.Contracts.InventoryController.Address).SetBehavior(ctx, economy.Contracts.Behaviors.PurchaseItem.Address, state)
		if err != nil {
			return fmt.Errorf("failed to set purchase_item behavior: %w", err)
		}

	default:
		return fmt.Errorf("unknown behavior: %s", behavior)
	}

	// FIXME this can cause update issues if multiple updates are happening at the same time
	_, err = s.store.UpdateGameEconomyContracts(ctx, userID, gameID, economyID, economy.Contracts)
	if err != nil {
		return fmt.Errorf("failed to update game economy: %w", err)
	}

	return nil
}
