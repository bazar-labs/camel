package service

import (
	"context"

	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/google/uuid"
)

func (s *Service) ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	return s.store.ListGames(ctx, userID)
}

func (s *Service) GetGame(ctx context.Context, userID uuid.UUID, gameID uuid.UUID) (*domain.Game, error) {
	return s.store.GetGame(ctx, userID, gameID)
}

func (s *Service) CreateGame(ctx context.Context, userID uuid.UUID, name string) (*domain.Game, error) {
	return s.store.CreateGame(ctx, userID, name)
}
