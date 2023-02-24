package controllers

import (
	"github.com/cruxia/go/initializers"
	"github.com/cruxia/go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}




func UpdateUser(c *fiber.Ctx) error {
	
	type updateUser struct {
		Role    string `gorm:"not null"`
	}
	var user models.User

	// Read the param noteId
	id := c.Params("userId")

	// Find the note with the given Id
	initializers.DB.Find(&user, "id = ?", id)

	// If no such note present return an error
	if *user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}
	// Store the body containing the updated data and return error if encountered
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the note
	*user.Role = updateUserData.Role

	// Save the Changes
	initializers.DB.Save(&user)

	// Return the updated note
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": user})
}
