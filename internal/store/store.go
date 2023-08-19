package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/bazar-labs/camel/internal/domain"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) ListGame(ctx context.Context, userID int64) ([]domain.Game, error) {
	query := `
		SELECT id, name
		FROM games
		WHERE user_id = ?
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying games: %w", err)
	}
	defer rows.Close()

	gg := make([]domain.Game, 0)
	for rows.Next() {
		var g domain.Game
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		gg = append(gg, g)
	}

	return gg, nil
}

func (s *Store) GetGame(ctx context.Context, userID, gameID int64) (*domain.Game, error) {
	query := `
		SELECT id, name
		FROM games
		WHERE user_id = ? AND id = ?
	`

	var g domain.Game
	if err := s.db.QueryRowContext(ctx, query, userID, gameID).Scan(&g.ID, &g.Name); err != nil {
		return nil, fmt.Errorf("error querying game: %w", err)
	}

	return &g, nil
}

func (s *Store) CreateGame(ctx context.Context, userID int64, name string) (*domain.Game, error) {
	statment := `
		INSERT INTO games (user_id, name)
		VALUES (?, ?)
	`

	res, err := s.db.ExecContext(ctx, statment, userID, name)
	if err != nil {
		return nil, fmt.Errorf("error creating game: %w", err)
	}

	gameID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %w", err)
	}

	return s.GetGame(ctx, userID, gameID)
}

func (s *Store) ListGameEconomy(ctx context.Context, userID, gameID int64) ([]domain.GameEconomy, error) {
	query := `
		SELECT ge.id, ge.game_id, ge.chain_network, ge.contracts
		FROM game_economies AS ge
		INNER JOIN games AS g ON ge.game_id = g.id
		WHERE g.user_id = ? AND g.id = ?
	`

	rows, err := s.db.QueryContext(ctx, query, userID, gameID)
	if err != nil {
		return nil, fmt.Errorf("error querying game economies: %w", err)
	}
	defer rows.Close()

	gg := make([]domain.GameEconomy, 0)
	for rows.Next() {
		var g domain.GameEconomy
		var contracts json.RawMessage
		if err := rows.Scan(&g.ID, &g.GameID, &g.ChainNetwork, &contracts); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		if err := json.Unmarshal(contracts, &g.Contracts); err != nil {
			return nil, fmt.Errorf("error unmarshaling contracts: %w", err)
		}
		gg = append(gg, g)
	}

	return gg, nil
}

func (s *Store) GetGameEconomy(ctx context.Context, userID, gameID, economyID int64) (*domain.GameEconomy, error) {
	query := `
		SELECT ge.id, ge.game_id, ge.chain_network, ge.contracts
		FROM game_economies AS ge
		INNER JOIN games AS g ON ge.game_id = g.id
		WHERE g.user_id = ? AND g.id = ? AND ge.id = ?
	`

	var g domain.GameEconomy
	var contracts json.RawMessage
	if err := s.db.QueryRowContext(ctx, query, userID, gameID, economyID).Scan(&g.ID, &g.GameID, &g.ChainNetwork, &contracts); err != nil {
		return nil, fmt.Errorf("error querying game economy: %w", err)
	}

	if err := json.Unmarshal(contracts, &g.Contracts); err != nil {
		return nil, fmt.Errorf("error unmarshaling contracts: %w", err)
	}

	return &g, nil
}

func (s *Store) CreateGameEconomy(ctx context.Context, userID, gameID int64, chainNetwork domain.ChainNetwork, contracts domain.GameEconomyContracts) (*domain.GameEconomy, error) {
	// TODO check game exists and belongs to user

	statment := `
		INSERT INTO game_economies (game_id, chain_network, contracts)
		VALUES (?, ?, ?)
	`

	aa, err := json.Marshal(contracts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling contracts: %w", err)
	}

	res, err := s.db.ExecContext(ctx, statment, gameID, chainNetwork, aa)
	if err != nil {
		return nil, fmt.Errorf("error creating game economy: %w", err)
	}

	economyID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %w", err)
	}

	return s.GetGameEconomy(ctx, userID, gameID, economyID)
}

func (s *Store) UpdateGameEconomyContracts(ctx context.Context, userID, gameID, economyID int64, contracts domain.GameEconomyContracts) (*domain.GameEconomyContracts, error) {
	statment := `
		UPDATE game_economies AS ge
		JOIN games AS g ON ge.game_id = g.id
		SET ge.contracts = ?
		WHERE g.user_id = ? AND g.id = ? AND ge.id = ?
	`

	data, err := json.Marshal(contracts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling contracts: %w", err)
	}

	if _, err := s.db.ExecContext(ctx, statment, data, userID, gameID, economyID); err != nil {
		return nil, fmt.Errorf("error updating game economy contracts: %w", err)
	}

	economy, err := s.GetGameEconomy(ctx, userID, gameID, economyID)
	if err != nil {
		return nil, fmt.Errorf("error getting game economy: %w", err)
	}

	return &economy.Contracts, nil
}

func (s *Store) GetMasterEconomyContracts(ctx context.Context, chainNetwork domain.ChainNetwork) (*domain.MasterEconomyContracts, error) {
	query := `
		SELECT contracts
		FROM master_economies
		WHERE chain_network = ?
	`

	var mc domain.MasterEconomyContracts
	var contracts json.RawMessage
	if err := s.db.QueryRowContext(ctx, query, chainNetwork).Scan(&contracts); err != nil {
		return nil, fmt.Errorf("error querying master contracts: %w", err)
	}

	if err := json.Unmarshal(contracts, &mc); err != nil {
		return nil, fmt.Errorf("error unmarshaling contracts: %w", err)
	}

	return &mc, nil
}
