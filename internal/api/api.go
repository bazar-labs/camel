package api

import (
	"context"
	"fmt"

	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type API struct {
	cfg    *config.API
	server *fiber.App
}

type handler struct {
	cfg     *config.API
	service IService
}

type middleware struct{}

type IService interface {
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
}

func New(cfg *config.API, service IService) *API {
	server := fiber.New()
	handler := &handler{cfg, service}
	middleware := &middleware{}

	server.Use(middleware.Logger())
	server.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	v1 := server.Group("/api/v1")
	v1.Get("/games", handler.ListGames)
	v1.Post("/upload", handler.Upload)

	return &API{cfg, server}
}

func (a *API) ListenAndServe() error {
	return a.server.Listen(fmt.Sprintf(":%d", a.cfg.ServerPort))
}