package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type Database struct {
	Name     string
	Password string
	Age      uint8
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Use(cors.New())

	app.Get("/api/v1/users", func(c *fiber.Ctx) {
		data := Database{"Eric", "Tacobell", 24}
		c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
		})
	})

	app.Static("/", "./frontend/build")

	fmt.Println("Server running on port 5010")
	app.Listen(5010)
}
