package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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
		return nil, fmt.Errorf("error querying games: %w", err)
	}
	defer rows.Close()

	// TODO CHECK
	games := make([]domain.Game, 0)
	for rows.Next() {
		var g domain.Game
		var contractAddresses json.RawMessage
		if err := rows.Scan(&g.ID, &g.Name, &contractAddresses); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		if err := json.Unmarshal(contractAddresses, &g.ContractAddresses); err != nil {
			return nil, fmt.Errorf("error unmarshalling contract addresses: %w", err)
		}
		games = append(games, g)
	}

	return games, nil
}

func (s *Store) GetGame(ctx context.Context, userID, gameID uuid.UUID) (*domain.Game, error) {
	query := `
		SELECT id, name, contract_addresses
		FROM games
		WHERE id = $1 AND user_id = $2
	`

	var g domain.Game
	var contractAddresses json.RawMessage
	if err := s.db.QueryRowContext(ctx, query, gameID, userID).Scan(&g.ID, &g.Name, &contractAddresses); err != nil {
		return nil, fmt.Errorf("error querying game: %w", err)
	}
	if err := json.Unmarshal(contractAddresses, &g.ContractAddresses); err != nil {
		return nil, fmt.Errorf("error unmarshalling contract addresses: %w", err)
	}

	return &g, nil
}

func (s *Store) CreateGame(ctx context.Context, userID uuid.UUID, name string) (*domain.Game, error) {
	query := `
		INSERT INTO games (user_id, name)
		VALUES ($1, $2)
		RETURNING id
	`

	var gameID uuid.UUID
	if err := s.db.QueryRowContext(ctx, query, userID, name).Scan(&gameID); err != nil {
		return nil, fmt.Errorf("error creating game: %w", err)
	}

	return s.GetGame(ctx, userID, gameID)
}

// TODO add userID
// TODO replace with GetGame?
func (s *Store) GetGameContractAddresses(ctx context.Context, gameID uuid.UUID) (*domain.GameContractAddresses, error) {
	query := `
		SELECT contract_addresses
		FROM games
		WHERE id = $1
	`

	var contractAddresses json.RawMessage
	if err := s.db.QueryRowContext(ctx, query, gameID).Scan(&contractAddresses); err != nil {
		return nil, fmt.Errorf("error querying game: %w", err)
	}

	var ca domain.GameContractAddresses
	if err := json.Unmarshal(contractAddresses, &ca); err != nil {
		return nil, fmt.Errorf("error unmarshalling contract addresses: %w", err)
	}

	return &ca, nil
}

// TODO add userID
// TODO replace with UpdateGame?
func (s *Store) UpdateGameContractAddresses(ctx context.Context, gameID uuid.UUID, addresses domain.GameContractAddresses) error {
	query := `
		UPDATE games
		SET contract_addresses = $1
		WHERE id = $2
	`

	contractAddresses, err := json.Marshal(addresses)
	if err != nil {
		return fmt.Errorf("error marshalling contract addresses: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, contractAddresses, gameID)
	if err != nil {
		return fmt.Errorf("error updating game: %w", err)
	}

	return nil
}
