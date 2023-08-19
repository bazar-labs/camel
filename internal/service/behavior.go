package service

import (
	"context"
	"fmt"

	"github.com/bazar-labs/turron/internal/domain"
)

func (s *Service) GetBehaviorState(ctx context.Context, userID, gameID, economyID int64, behavior domain.Behavior) (bool, error) {
	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return false, err
	}

	controller := s.blockchain.InventoryController(economy.Contracts.InventoryController)

	switch behavior {
	case domain.BehaviorPurchaseItemWithETH:
		return controller.IsBehaviorEnabled(economy.Contracts.Behaviors.PurchaseItemWithETH)
	default:
		return false, fmt.Errorf("unknown behavior: %s", behavior)
	}
}

func (s *Service) EnableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior domain.Behavior) error {
	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return err
	}

	controller := s.blockchain.InventoryController(economy.Contracts.InventoryController)

	switch behavior {
	case domain.BehaviorPurchaseItemWithETH:
		err := controller.EnableBehavior(ctx, economy.Contracts.Behaviors.PurchaseItemWithETH)
		if err != nil {
			return fmt.Errorf("failed to enable behavior: %w", err)
		}

	default:
		return fmt.Errorf("unknown behavior: %s", behavior)
	}

	return nil
}

func (s *Service) DisableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior domain.Behavior) error {
	economy, err := s.store.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return err
	}

	controller := s.blockchain.InventoryController(economy.Contracts.InventoryController)

	switch behavior {
	case domain.BehaviorPurchaseItemWithETH:
		err := controller.DisableBehavior(ctx, economy.Contracts.Behaviors.PurchaseItemWithETH)
		if err != nil {
			return fmt.Errorf("failed to disable behavior: %w", err)
		}

	default:
		return fmt.Errorf("unknown behavior: %s", behavior)
	}

	return nil
}
