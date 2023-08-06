package api

import (
	"context"
	"fmt"
	"math/big"
	"mime/multipart"

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
	// Game
	ListGames(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
	GetGame(ctx context.Context, userID uuid.UUID, gameID uuid.UUID) (*domain.Game, error)
	CreateGame(ctx context.Context, userID uuid.UUID, name string) (*domain.Game, error)

	// Game Economy
	DeployGameEconomy(ctx context.Context, gameID uuid.UUID) (domain.GameContractAddresses, error)

	// Inventory Registry
	GetItemDefinition(ctx context.Context, gameID uuid.UUID, itemDefID *big.Int) (*domain.ItemDefinition, error)
	CreateItemDefinition(ctx context.Context, gameID uuid.UUID, form *multipart.Form) (*domain.ItemDefinition, error)
}

func New(cfg *config.API, service IService) *API {
	server := fiber.New()
	handler := &handler{cfg, service}
	middleware := &middleware{}

	server.Use(middleware.Logger())
	server.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	v1 := server.Group("/api/v1")
	v1.Get("/games", handler.ListGames)
	v1.Get("/games/:id", handler.GetGame)
	v1.Post("/games", handler.CreateGame)
	v1.Post("/games/:id/economy", handler.DeployGameEconomy)
	v1.Get("/games/:id/economy/inventory/registry/:item_def_id", handler.GetItemDefinition)
	v1.Post("/games/:id/economy/inventory/registry", handler.CreateItemDefinition)

	return &API{cfg, server}
}

func (a *API) ListenAndServe() error {
	return a.server.Listen(fmt.Sprintf(":%d", a.cfg.ServerPort))
}
