package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	v1 "github.com/filipeandrade6/cooperagro/cmd/api/v1"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// =================================================================================
	// configuration

	// =================================================================================
	// database

	dataSourceName := "postgresql://postgres:postgres@localhost:5432/cooperagro"
	db, err := postgres.NewPostgresRepo(dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	// =================================================================================
	// API service

	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	v1.RegisterHandlers(e, db)

	// =================================================================================
	// graceful shutdown

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
