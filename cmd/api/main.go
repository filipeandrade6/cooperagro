package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filipeandrade6/cooperativa-go/foundation/logger"
	"github.com/filipeandrade6/cooperativa-go/handlers"

	"github.com/ardanlabs/conf/v3"
	_ "go.uber.org/automaxprocs"
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
		log.Error("startup", zap.Error(err))
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.Logger) error {
	log.Info("initializing API support")

	// Create a configuration struct with default values.
	cfg := struct {
		ReadTimeout     time.Duration `conf:"default:5s"`
		WriteTimeout    time.Duration `conf:"default:10s"`
		IdleTimeout     time.Duration `conf:"default:120s"`
		ShutdownTimeout time.Duration `conf:"default:20s"`
		APIHost         string        `conf:"default:0.0.0.0:8080,env:API_HOST"`
	}{}

	// Parse the configuration.
	const prefix = "COOPERAGRO"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// Make a channel to listen for INTERRUPT/TERMINATE signal from user.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct the API.
	api := handlers.NewAPI(log)

	// Construct the server to service the requests.
	srv := &http.Server{
		Addr:         cfg.APIHost,
		Handler:      api,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		ErrorLog:     zap.NewStdLog(log),
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)
	go func() {
		log.Info("API router started", zap.String("host", cfg.APIHost))
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info("shutdown started", zap.String("signal", sig.String()))
		defer log.Info("shutdown complete", zap.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
