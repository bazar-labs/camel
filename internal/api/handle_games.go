package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var userID = 1
var gameID = 1

func (h *handler) ListGames(c *fiber.Ctx) error {
	games, err := h.service.ListGames(c.Context(), userID)
	if err != nil {
		return fmt.Errorf("failed to list games: %w", err)
	}

	return c.JSON(games)
}

func (h *handler) GetGame(c *fiber.Ctx) error {
	game, err := h.service.GetGame(c.Context(), userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to get game: %w", err)
	}

	return c.JSON(game)
}

func (h *handler) CreateGame(c *fiber.Ctx) error {
	game, err := h.service.CreateGame(c.Context(), userID, "My Game")
	if err != nil {
		return fmt.Errorf("failed to create game: %w", err)
	}

	return c.JSON(game)
}
