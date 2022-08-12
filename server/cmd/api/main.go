package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	v1 "github.com/filipeandrade6/cooperagro/cmd/api/v1"
	"github.com/filipeandrade6/cooperagro/infra/config"
	"github.com/filipeandrade6/cooperagro/infra/logger"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New("COOPERAGRO")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("initializing", "error", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// =================================================================================
	// configuration

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("getting config: %w", err)
	}

	// =================================================================================
	// database

	db, err := postgres.NewPostgresRepo(postgres.Config{
		Host:     cfg.DB.Host,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		TLS:      cfg.DB.TLS,
	})
	if err != nil {
		return fmt.Errorf("database: %w", err)
	}

	// =================================================================================
	// API service

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Desugar().Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	v1.RegisterHandlers(e, db)

	// =================================================================================
	// graceful shutdown

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server") // TODO não esta logando
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err) // TODO não esta logando
	}

	return nil
}
