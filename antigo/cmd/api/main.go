package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/filipeandrade6/cooperagro/assets"

	"github.com/filipeandrade6/cooperagro/cmd/api/middlewares"
	"github.com/filipeandrade6/cooperagro/cmd/api/productsctrl"
	"github.com/filipeandrade6/cooperagro/cmd/api/usersctrl"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/filipeandrade6/cooperagro/domain/products"
	"github.com/filipeandrade6/cooperagro/domain/users"

	"github.com/filipeandrade6/cooperagro/adapters/log"
	"github.com/filipeandrade6/cooperagro/adapters/log/jsonlogs"
	"github.com/filipeandrade6/cooperagro/helpers/env"

	"github.com/filipeandrade6/cooperagro/adapters/repo"
	"github.com/filipeandrade6/cooperagro/adapters/repo/pgrepo"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	// ========================================================================
	// Configuration
	port := env.GetString("PORT", "80")
	logLevel := env.GetString("LOG_LEVEL", "INFO")
	// foursquareBaseURL := env.MustGetString("FOURSQUARE_BASE_URL")
	// foursquareClientID := env.MustGetString("FOURSQUARE_CLIENT_ID")
	// foursquareSecret := env.MustGetString("FOURSQUARE_SECRET")
	// redisURL := env.GetString("REDIS_URL", "")
	// redisPassword := env.GetString("REDIS_PASSWORD", "")
	dbURL := env.MustGetString("DATABASE_URL")

	// Dependency Injection goes here:
	logger := jsonlogs.New(logLevel, domain.GetCtxValues)

	// restClient := http.New(30 * time.Second)

	// var cacheClient cache.Provider
	// if redisURL != "" {
	// 	cacheClient = redis.New(redisURL, redisPassword, 24*time.Hour)
	// } else {
	// 	cacheClient = memorycache.New(24*time.Hour, 10*time.Minute)
	// }

	// venuesService := venues.NewService(
	// 	logger,
	// 	restClient,
	// 	cacheClient,
	// 	foursquareBaseURL,
	// 	foursquareClientID,
	// 	foursquareSecret,
	// )

	// // The controllers handle HTTP stuff so the services can be kept as simple as possible
	// // only working on top of the domain language, i.e. types and interfaces from the domain package
	// venuesController := venuesctrl.NewController(venuesService)

	var repository repo.Repo
	repository, err := pgrepo.New(ctx, dbURL)
	if err != nil {
		logger.Fatal(ctx, "unable to start database", log.Body{
			"db_url": dbURL,
			"error":  err.Error(),
		})
	}

	usersService := users.NewService(logger, repository)
	usersController := usersctrl.NewController(usersService)

	productsService := products.NewService(logger, repository)
	productsController := productsctrl.NewController(productsService)

	// Any framework you need for serving HTTP or GRPC goes in the main package,
	//
	// It should be kept here because the main package is the only one that is allowed
	// to depend on anything, and also because this logic is unique to this endpoint,
	// so you won't reuse it anywhere else.
	app := fiber.New()

	app.Use(middlewares.HandleRequestID())
	app.Use(middlewares.HandleError(logger))
	app.Use(middlewares.RequestLogger(logger))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/products/:id", productsController.GetProduct)

	app.Post("/users", usersController.UpsertUser)
	app.Get("/users/:id", usersController.GetUser)

	// app.Get("/venues/:latitude,:longitude", venuesController.GetVenuesByCoordinates)
	// app.Get("/venues/details/:id", venuesController.GetDetails)

	// Just an example on how to serve html templates using the embed library
	// and explicit arguments with a "builder function":
	app.Get("/example-html", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return assets.WriteExamplePage(c, "username", "user address", 42)
	})

	logger.Info(ctx, "server-starting-up", log.Body{
		"port": port,
	})
	if err := app.Listen(":" + port); err != nil {
		logger.Error(ctx, "server-stopped-with-an-error", log.Body{
			"error": err.Error(),
		})
	}
}