package service

import (
	"context"

	"github.com/el-goblino-foundation/turron/domain"
	"github.com/google/uuid"
)

type Service struct {
	store IStore
}

type IStore interface {
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
}

func New(store IStore) *Service {
	return &Service{store}
}

func (s *Service) ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	return s.store.ListGames(ctx, userID)
}
