package store

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/el-goblino-foundation/turron/internal/domain"
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
		SELECT id, name, contract_addresses
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
		var contractAddresses json.RawMessage
		if err := rows.Scan(&g.ID, &g.Name, &contractAddresses); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(contractAddresses, &g.ContractAddresses); err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

func (s *Store) CreateGame(ctx context.Context, userID uuid.UUID, name string) (uuid.UUID, error) {
	query := `
		INSERT INTO games (user_id, name)
		VALUES ($1, $2)
		RETURNING id
	`

	var gameID uuid.UUID
	if err := s.db.QueryRowContext(ctx, query, userID, name).Scan(&gameID); err != nil {
		return uuid.Nil, err
	}

	return gameID, nil
}

func (s *Store) GetGameContractAddresses(ctx context.Context, gameID uuid.UUID) (*domain.GameContractAddresses, error) {
	query := `
		SELECT contract_addresses
		FROM games
		WHERE id = $1
	`

	var contractAddresses json.RawMessage
	if err := s.db.QueryRowContext(ctx, query, gameID).Scan(&contractAddresses); err != nil {
		return nil, err
	}

	var ca domain.GameContractAddresses
	if err := json.Unmarshal(contractAddresses, &ca); err != nil {
		return nil, err
	}

	return &ca, nil
}

func (s *Store) UpdateGameContractAddresses(ctx context.Context, gameID uuid.UUID, addresses domain.GameContractAddresses) error {
	query := `
		UPDATE games
		SET contract_addresses = $1
		WHERE id = $2
	`

	contractAddresses, err := json.Marshal(addresses)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, contractAddresses, gameID)
	if err != nil {
		return err
	}

	return nil
}
