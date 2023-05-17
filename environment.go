package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type env uint

const (
	development              env = 1 // DO NOT USE THIS IN PRODUCTION
	testing                  env = 2 // DO NOT USE THIS IN PRODUCTION
	production               env = 3
	developmentWithoutTokens env = 4 // DO NOT USE THIS IN PRODUCTION
)

func initEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file:", err)
	}
}

func EnvironmentsManager(app *fiber.App, env env) {

	if env > 0 && env < 5 {
		switch env {
		case development:
			fmt.Println("Running in development mode")
			initEnvs()
			defaultInitConf(app, true)
		case testing:
			fmt.Println("Running in testing mode")
			initEnvs()
			defaultInitConf(app, true)
		case production:
			fmt.Println("Running in production mode")
			defaultInitConf(app, true)
		case developmentWithoutTokens:
			fmt.Println("Running in development mode without tokens")
			initEnvs()
			defaultInitConf(app, false)
		}
	} else {
		fmt.Print("Invalid environment configuration")
	}

}
