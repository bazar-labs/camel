package service

import (
	"context"

	"github.com/el-goblino-foundation/turron/internal/domain"
)

func (s *Service) ListGames(ctx context.Context, userID int) ([]domain.Game, error) {
	return s.store.ListGames(ctx, userID)
}

func (s *Service) GetGame(ctx context.Context, userID int, gameID int) (*domain.Game, error) {
	return s.store.GetGame(ctx, userID, gameID)
}

func (s *Service) CreateGame(ctx context.Context, userID int, name string) (*domain.Game, error) {
	return s.store.CreateGame(ctx, userID, name)
}
