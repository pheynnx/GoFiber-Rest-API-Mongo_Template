package main

import (
	"fmt"
	"os"

	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/database"
	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/models"
	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviromental variables from config
	godotenv.Load("./config/config.env")

	// Create new fiber instance
	app := fiber.New(fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})

	// Connect to database
	database.Connect()
	defer database.Cancel()
	defer database.Client.Disconnect(database.Ctx)

	// Create model schemas
	models.CreateUserSchema()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	routes.UserRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
