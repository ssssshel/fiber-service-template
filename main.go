package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ssssshel/fiber-service-template/src/middlewares"
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
}
