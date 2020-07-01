package controllers

import (
	"ericarthurc/fiberAPI/database"
	"ericarthurc/fiberAPI/models"

	"github.com/gofiber/fiber"
)

// GetUsers get all users
func GetUsers(c *fiber.Ctx) {
	db := database.DB
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.Status(500)
		return
	}
	c.JSON(&fiber.Map{
		"success": true,
		"data":    users,
	})
}

// GetUser get user by id
func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var user models.User
	if err := db.Find(&user, id).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "data": err})
		return

	}
	c.JSON(fiber.Map{"success": true, "data": user})
}

// CreateUser create user
func CreateUser(c *fiber.Ctx) {
	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "data": err})
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "data": err})
		return
	}
	c.JSON(fiber.Map{"success": true, "data": user})

}

func UpdateUser(c *fiber.Ctx) {
	db
}

// DeleteUser deletes user by id
func DeleteUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "data": err})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "data": err})
		return
	}
	c.JSON(fiber.Map{"success": true})
}
