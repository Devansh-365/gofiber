package main

import (
	configs "github.com/Devansh-365/gofiber/configs"
	"github.com/Devansh-365/gofiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
