package routes

import (
	"github.com/bdn/jeker/controllers"
	"github.com/gofiber/fiber"
)

func UserRoutes(app *fiber.App) {
	app.Get("/api/v1/user/health", func(c *fiber.Ctx) {
		c.SendString("MESSAGE FROM USER ROUTE")
	})
	app.Post("/api/v1/user/register", controllers.UserRegister)
	app.Post("/api/v1/user/login", controllers.UserLogin)
}
