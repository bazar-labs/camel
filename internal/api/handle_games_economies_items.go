package api

import (
	"math/big"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) ListItem(c *fiber.Ctx) error {
	items, err := h.service.ListItem(c.Context(), userID, gameID, economyID)
	if err != nil {
		return err
	}

	return c.JSON(items)
}

func (h *handler) GetItem(c *fiber.Ctx) error {
	type params struct {
		ItemDefID uint `params:"item_def_id"`
	}

	var p params
	if err := c.ParamsParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query params"})
	}

	item, err := h.service.GetItem(c.Context(), userID, gameID, economyID, big.NewInt(int64(p.ItemDefID)))
	if err != nil {
		return err
	}

	return c.JSON(item)
}

func (h *handler) CreateItem(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	item, err := h.service.CreateItem(c.Context(), userID, gameID, economyID, form)
	if err != nil {
		return err
	}

	return c.JSON(item)
}
