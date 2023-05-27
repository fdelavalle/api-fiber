package main

import (
	"log"

	"github.com/fdelavalle/api-fiber/configs"
	"github.com/fdelavalle/api-fiber/database"
	"github.com/fdelavalle/api-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("api/books", routes.AllBooks)
	app.Post("api/books", routes.AddBook)
	app.Get("api/books/:id", routes.GetBook)
	app.Put("api/books/:id", routes.UpdateBook)
	app.Delete("api/books/:id", routes.DeleteBook)
}

func main() {
	configs.InitEnvConfigs()
	database.ConnectDb(configs.EnvConfigs.DatabaseUrl)
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
