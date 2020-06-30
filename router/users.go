package router

import (
	"ericarthurc/fiberAPI/controllers"

	"github.com/gofiber/fiber"
)

// UserRoutes routes for /api/v1/users
func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
}
