package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ssssshel/fiber-service-template/src/middlewares"
	"github.com/ssssshel/fiber-service-template/src/shared/config"
)

func main() {
	app := fiber.New()

	EnvironmentsManager(app, development)
}

func defaultInitConf(app *fiber.App, tokenization bool) {

	app.Use(logger.New())
	app.Use(cors.New())

	if tokenization {
		app.Use(middlewares.AccessTokenChecker)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MS TEMPLATE")
	})

	fmt.Printf("Server running on port %s\n", config.Port())
	app.Listen(":" + config.Port())
}
