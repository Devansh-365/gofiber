package routes

import (
	"github.com/Devansh-365/gofiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/users/:userId", controllers.GetAUser)
	app.Get("/users", controllers.GetAllUsers)
}
