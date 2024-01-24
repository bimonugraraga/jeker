package routes

import (
	"github.com/bdn/jeker/controllers"
	"github.com/bdn/jeker/middlewares"
	"github.com/gofiber/fiber"
)

func ProfileRoutes(app *fiber.App) {
	app.Get("/api/v1/profile/health", func(c *fiber.Ctx) {
		c.SendString("MESSAGE FROM USER PROFILE")
	})

	app.Post("/api/v1/profile", middlewares.UserAuthentication, controllers.CreateUpdateProfile)
	app.Get("/api/v1/profile", controllers.GetListOfProfile)
	app.Get("/api/v1/profile/:user_id", controllers.GetOneProfile)
}
