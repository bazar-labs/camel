package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (m *middleware) Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		code := c.Response().StatusCode()

		sublog := log.With().
			Int("status", code).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("latency", time.Since(start).Round(time.Millisecond).String()).
			Str("ip", c.IP()).
			Str("user-agent", c.Get(fiber.HeaderUserAgent)).
			Logger()

		switch {
		case http.StatusInternalServerError <= code || err != nil:
			sublog.Error().RawJSON("response", c.Response().Body()).Err(err).Msg("internal server error")
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})

		case fiber.StatusBadRequest <= code && code < fiber.StatusInternalServerError:
			sublog.Warn().RawJSON("response", c.Response().Body()).Msg("bad request")
			return nil

		default:
			sublog.Info().Msg("succesful request")
			return nil
		}
	}
}