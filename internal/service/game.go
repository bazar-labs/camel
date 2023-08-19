package service

import (
	"context"

	"github.com/bazar-labs/camel/internal/domain"
)

func (s *Service) ListGame(ctx context.Context, userID int64) ([]domain.Game, error) {
	return s.store.ListGame(ctx, userID)
}

func (s *Service) GetGame(ctx context.Context, userID int64, gameID int64) (*domain.Game, error) {
	return s.store.GetGame(ctx, userID, gameID)
}

func (s *Service) CreateGame(ctx context.Context, userID int64, name string) (*domain.Game, error) {
	return s.store.CreateGame(ctx, userID, name)
}
