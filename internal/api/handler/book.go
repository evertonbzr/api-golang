package handler

import (
	"strconv"

	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookRepo *repository.BookRepository
}

func NewBookHandler() *BookHandler {
	return &BookHandler{
		BookRepo: repository.NewBookRepository(),
	}
}

func (h *BookHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data types.CreateBookRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		book := model.Book{
			Title:       data.Title,
			Description: data.Description,
			Author:      data.Author,
		}

		if err := h.BookRepo.Create(
			[]model.Book{book},
		); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"message": "Create book successfully",
			"book":    book,
		})
	}
}

func (h *BookHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		books, err := h.BookRepo.List()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"books": books,
		})
	}
}

func (h *BookHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookId := c.Params("id")

		var data types.UpdateBookRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		id64, err := strconv.Atoi(bookId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}
		id := uint(id64)

		book, err := h.BookRepo.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "book not found",
			})
		}

		if data.Title != "" {
			book.Title = data.Title
		}
		if data.Description != "" {
			book.Description = data.Description
		}
		if data.Status != "" {
			book.Status = data.Status
		}
		if data.Author != "" {
			book.Author = data.Author
		}

		if err := h.BookRepo.Update(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update",
			})
		}

		return c.JSON(fiber.Map{
			"message": "book updated successfully",
			"book":    book,
		})
	}
}
