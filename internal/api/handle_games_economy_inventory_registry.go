package api

import (
	"math/big"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetItemDefinition(c *fiber.Ctx) error {
	type params struct {
		ItemDefID uint `param:"item_def_id"`
	}

	var p params
	if err := c.QueryParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query params"})
	}

	item, err := h.service.GetItemDefinition(c.Context(), gameID, big.NewInt(int64(p.ItemDefID)))
	if err != nil {
		return err
	}

	return c.JSON(item)
}

func (h *handler) CreateItemDefinition(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	item, err := h.service.CreateItemDefinition(c.Context(), gameID, form)
	if err != nil {
		return err
	}

	return c.JSON(item)
}
