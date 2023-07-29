package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var userID = uuid.MustParse("c1914084-f4c8-4d38-a1f1-2836cab7bdd5")

func (h *handler) ListGames(c *fiber.Ctx) error {
	games, err := h.service.ListGames(c.Context(), userID)
	if err != nil {
		return fmt.Errorf("failed to list games: %v", err)
	}

	return c.JSON(games)
}
