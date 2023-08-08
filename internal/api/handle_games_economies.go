package api

import (
	"fmt"

	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) ListGameEconomy(c *fiber.Ctx) error {
	gameEconomy, err := h.service.ListGameEconomy(c.Context(), userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to list game economy: %w", err)
	}

	return c.JSON(gameEconomy)
}

func (h *handler) GetGameEconomy(c *fiber.Ctx) error {
	gameEconomy, err := h.service.GetGameEconomy(c.Context(), userID, gameID, economyID)
	if err != nil {
		return fmt.Errorf("failed to get game economy: %w", err)
	}

	return c.JSON(gameEconomy)
}

func (h *handler) CreateGameEconomy(c *fiber.Ctx) error {
	// FIXME userID, gameID, domain.Localhost
	addresses, err := h.service.CreateGameEconomy(c.Context(), userID, gameID, domain.Localhost)
	if err != nil {
		return fmt.Errorf("failed to deploy game economy: %w", err)
	}

	return c.JSON(addresses)
}
