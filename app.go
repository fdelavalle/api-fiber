package main

import (
	"log"

	"github.com/fdelavalle/api-fiber/configs"
	"github.com/fdelavalle/api-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	configs.InitEnvConfigs()
	database.ConnectDb(configs.EnvConfigs.DatabaseUrl)
	app := fiber.New()

	/* setUpRoutes(app) */

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
