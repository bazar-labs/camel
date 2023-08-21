package api

import "github.com/gofiber/fiber/v2"

func (h *handler) EnableBehavior(c *fiber.Ctx) error {
	err := h.service.EnableBehavior(c.Context(), userID, gameID, economyID, behavior)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) DisableBehavior(c *fiber.Ctx) error {
	err := h.service.DisableBehavior(c.Context(), userID, gameID, economyID, behavior)
	if err != nil {
		return nil
	}

	return c.SendStatus(fiber.StatusOK)
}
