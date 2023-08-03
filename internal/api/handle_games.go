package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var userID = uuid.MustParse("c1914084-f4c8-4d38-a1f1-2836cab7bdd5")
var gameID = uuid.MustParse("3918ba29-75dd-40e9-aa5a-35d83ee75d40")

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
