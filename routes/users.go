package routes

import (
	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/controllers"
	"github.com/gofiber/fiber/v2"
)

// UserRoutes | @desc: "/api/v1/users" api route
func UserRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
	users.Patch("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}
