package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	cfg "github.com/revaldidwioctavian/go-fiber-hex-arch/internal/adapter/config"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/adapter/handler"
	mylogger "github.com/revaldidwioctavian/go-fiber-hex-arch/internal/adapter/logger"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/service"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/storage/mongodb"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/storage/mongodb/repository"
)

func main() {
	config, err := cfg.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	mylogger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	db, _ := mongodb.NewDB(config.DB)

	repo := repository.NewProductRepository(db)
	service := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(service)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}]:${port} | ${status} | ${method} ${path} | ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))

	handler.NewRouter(app, *productHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Wello Horld!")
	})

	listenPort := fmt.Sprintf(":%s", config.HTTP.Port)
	app.Listen(listenPort)
}
