package api

import (
	"context"
	"fmt"
	"math/big"
	"mime/multipart"

	"github.com/bazar-labs/turron/internal/config"
	"github.com/bazar-labs/turron/internal/domain"
	"github.com/gofiber/fiber/v2"
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
	ListGame(ctx context.Context, userID int64) ([]domain.Game, error)
	GetGame(ctx context.Context, userID int64, gameID int64) (*domain.Game, error)
	CreateGame(ctx context.Context, userID int64, name string) (*domain.Game, error)

	// Game Economy
	ListGameEconomy(ctx context.Context, userID, gameID int64) ([]domain.GameEconomy, error)
	GetGameEconomy(ctx context.Context, userID, gameID, economyID int64) (*domain.GameEconomy, error)
	CreateGameEconomy(ctx context.Context, userID, gameID int64, chainNetwork domain.ChainNetwork) (*domain.GameEconomy, error)

	// Game Economy Item
	ListItem(ctx context.Context, userID, gameID, economyID int64) ([]domain.Item, error)
	GetItem(ctx context.Context, userID, gameID, economyID int64, itemDefID *big.Int) (*domain.Item, error)
	CreateItem(ctx context.Context, userID, gameID, economyID int64, form *multipart.Form) (*domain.Item, error)

	// Game Economy Behavior
	EnableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior string) error
	DisableBehavior(ctx context.Context, userID, gameID, economyID int64, behavior string) error
}

func New(cfg *config.API, service IService) *API {
	server := fiber.New()
	handler := &handler{cfg, service}
	middleware := &middleware{}
	setup(server, handler, middleware)
	return &API{cfg, server}
}

func (a *API) ListenAndServe() error {
	return a.server.Listen(fmt.Sprintf(":%d", a.cfg.ServerPort))
}

func setup(server *fiber.App, handler *handler, middleware *middleware) {
	server.Use(middleware.Logger())
	server.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	v1 := server.Group("/api/v1")

	// Game
	v1.Get("/games", handler.ListGame)
	v1.Get("/games/:gid", handler.GetGame)
	v1.Post("/games", handler.CreateGame)

	// Game Economy
	v1.Get("/games/:gid/economies", handler.ListGameEconomy)
	v1.Get("/games/:gid/economies/:eid", handler.GetGameEconomy)
	v1.Post("/games/:gid/economies", handler.CreateGameEconomy)

	// Game Economy Item
	v1.Get("/games/:gid/economies/:eid/items", handler.ListItem)
	v1.Get("/games/:gid/economies/:eid/items/:iid", handler.GetItem)
	v1.Post("/games/:gid/economies/:eid/items", handler.CreateItem)

	// Game Economy Behavior
	v1.Post("/games/:gid/economies/:eid/behavior/:type/enable", handler.EnableBehavior)
	v1.Post("/games/:gid/economies/:eid/behavior/:type/disable", handler.DisableBehavior)
}
