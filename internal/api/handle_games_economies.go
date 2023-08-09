package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) ListGameEconomy(c *fiber.Ctx) error {
	economy, err := h.service.ListGameEconomy(c.Context(), userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to list game economy: %w", err)
	}

	return c.JSON(economy)
}

func (h *handler) GetGameEconomy(c *fiber.Ctx) error {
	economy, err := h.service.GetGameEconomy(c.Context(), userID, gameID, economyID)
	if err != nil {
		return fmt.Errorf("failed to get game economy: %w", err)
	}

	return c.JSON(economy)
}

func (h *handler) CreateGameEconomy(c *fiber.Ctx) error {
	economy, err := h.service.CreateGameEconomy(c.Context(), userID, gameID, chainNetwork)
	if err != nil {
		return fmt.Errorf("failed to create game economy: %w", err)
	}

	return c.JSON(economy)
}
