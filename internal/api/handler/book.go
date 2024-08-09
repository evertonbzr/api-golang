package handler

import (
	"strconv"

	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct {
	Service *service.BookService
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		Service: service.NewBookService(db),
	}
}

func (h *BookHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.CreateOrUpdateBookRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		book := model.Book{
			Title:       data.Title,
			Description: data.Description,
			Status:      data.Status,
			Author:      data.Author,
		}

		if err := h.Service.Create(
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
		books, err := h.Service.List()

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

		var data types.CreateOrUpdateBookRequest

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

		book, err := h.Service.GetByID(id)
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

		if err := h.Service.Update(book); err != nil {
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
