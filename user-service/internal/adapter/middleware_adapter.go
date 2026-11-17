package adapter

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/config"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/handler/response"
)

type MiddlewareAdapterInterface interface {
	CheckToken() echo.MiddlewareFunc
}

type middlewareAdapter struct {
	cfg *config.Config
}

func NewMiddlewareAdapter(cfg *config.Config) MiddlewareAdapterInterface {
	return &middlewareAdapter{
		cfg: cfg,
	}
}

func (m *middlewareAdapter) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			respErr := response.Response{}
			redisConn := config.NewRedisClient()
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				log.Error().Msg("authHeader is empty")
				respErr.Message = "missing or invalid token"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			getSession, err := redisConn.HGetAll(c.Request().Context(), tokenStr).Result()
			if err != nil || len(getSession) == 0 {
				log.Error().Msg("session not found")
				respErr.Message = "missing or invalid token"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			c.Set("user", getSession)
			return next(c)
		}
	}
}
