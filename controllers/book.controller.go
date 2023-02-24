package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cruxia/go/models"
	"github.com/cruxia/go/initializers"
	
)


func GetBook(c *fiber.Ctx) error {
	var books []models.Book

	// find all notes in the database
	
	initializers.DB.Find(&books)

	// If no note is present return an error
	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": books})
}

func CreateNote(c *fiber.Ctx) error {
	book := new(models.Book)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the Note and return error if encountered
	err = initializers.DB.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": book})
}