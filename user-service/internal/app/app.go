package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/config"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/handler"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/repository"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/service"
	"github.com/terenjit/ecommerce-veggie/user-service/utils/validator"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return
	}

	userRepo := repository.NewUserRepository(db.DB)
	jwtService := service.NewJwtService(cfg)
	userService := service.NewUserService(userRepo, cfg, jwtService)

	e := echo.New()
	e.Use(middleware.CORS())

	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	handler.NewUserHandler(e, userService, cfg)

	go func() {
		port := cfg.App.AppPort
		if port == "" {
			port = os.Getenv("APP_PORT")
		}
		if port == "" {
			port = "8080"
		}
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
		if err := e.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

}
