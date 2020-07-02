package main

import (
	"fmt"
	"os"

	"ericarthurc/fiberAPI/database"
	"ericarthurc/fiberAPI/routers"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	godotenv.Load("./config/config.env")

	// Connect to database
	database.ConnectDB()

	app := fiber.New(&fiber.Settings{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})

	// Middleware
	app.Use(middleware.Logger())
	app.Use(cors.New())

	// User route
	routers.UserRoutes(app)

	// Serve static frontend build
	// app.Static("/", "./frontend/build")

	fmt.Printf("Server running on port %v\n", os.Getenv("PORT"))
	app.Listen(os.Getenv("PORT"))

	defer database.DB.Close()
}
