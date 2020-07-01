package main

import (
	"fmt"
	"os"

	"ericarthurc/fiberAPI/database"
	"ericarthurc/fiberAPI/router"

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

	router.UserRoutes(app)

	app.Static("/", "./frontend/build")

	fmt.Printf("Server running on port %v\n", os.Getenv("PORT"))
	app.Listen(os.Getenv("PORT"))

	defer database.DB.Close()
}
