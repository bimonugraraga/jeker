package routes

import (
	"github.com/bdn/jeker/controllers"
	"github.com/bdn/jeker/middlewares"
	"github.com/gofiber/fiber"
)

func BookRoutes(app *fiber.App) {
	app.Get("/api/v1/book/health", func(c *fiber.Ctx) {
		c.SendString("MESSAGE FROM BOOK ROUTE")
	})
	app.Post("/api/v1/book", middlewares.UserAuthentication, controllers.CreateBook)
	app.Put("/api/v1/book/:book_id", middlewares.UserAuthentication, controllers.UpdateBook)
	app.Get("/api/v1/book/:book_id")
}
