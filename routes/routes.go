package routes

import (
	"github.com/fdelavalle/api-fiber/database"
	"github.com/fdelavalle/api-fiber/models"
	"github.com/gofiber/fiber/v2"
)

// AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&book)

	return c.Status(200).JSON(book)
}

// Update
func UpdateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	bookId := c.Params("id")

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Model(&book).Where("id = ?", bookId).Updates(book)
	return c.Status(200).JSON(book)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DB.Db.Find(&books)

	return c.Status(200).JSON(books)
}

// Book
func GetBook(c *fiber.Ctx) error {
	book := []models.Book{}
	bookId := c.Params("id")

	database.DB.Db.Where("id = ?", bookId).Find(&book)

	return c.Status(200).JSON(book)
}

// Delete
func DeleteBook(c *fiber.Ctx) error {
	book := new(models.Book)
	bookId := c.Params("id")

	database.DB.Db.Delete(&book, bookId)

	return c.Status(200).JSON("deleted")
}
