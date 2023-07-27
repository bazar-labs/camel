package store

import (
	"context"
	"database/sql"

	"github.com/el-goblino-foundation/turron/domain"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	query := `
		SELECT id, user_id, name
		FROM games
		WHERE user_id = $1
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var g domain.Game
		if err := rows.Scan(&g.ID, &g.UserID, &g.Name); err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}
