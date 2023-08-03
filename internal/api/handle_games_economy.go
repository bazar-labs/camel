package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) DeployGameEconomy(c *fiber.Ctx) error {
	addresses, err := h.service.DeployGameEconomy(c.Context(), gameID)
	if err != nil {
		return fmt.Errorf("failed to deploy game economy: %w", err)
	}

	return c.JSON(addresses)
}
