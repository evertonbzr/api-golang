package handler

import (
	"time"

	"github.com/evertonbzr/api-golang/internal/api/types"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type BorrowingHandler struct {
	BorrowingRepo *repository.BorrowingRepository
	BookRepo      *repository.BookRepository
	UserRepo      *repository.UserRepository
}

func NewBorrowingHandler() *BorrowingHandler {
	return &BorrowingHandler{
		BorrowingRepo: repository.NewBorrowingRepository(),
		BookRepo:      repository.NewBookRepository(),
		UserRepo:      repository.NewUserRepository(),
	}
}

func (h *BorrowingHandler) Set() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data types.SetBorrowingRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		var borrowing model.Borrowing
		var action string

		if data.UserID != 0 {
			user, err := h.UserRepo.GetUserById(data.UserID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "User not found",
				})
			}

			if user.Role != "user" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "User is not allowed to borrow",
				})
			}

			borrowing.UserID = user.ID
		}

		if data.BookID != 0 {
			book, err := h.BookRepo.GetByID(data.BookID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Book not found",
				})
			}

			if book.Status == "borrowed" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Book is already borrowed",
				})
			}

			borrowing.BookID = book.ID
		}

		if data.ID == 0 {
			action = "create"
			borrowing.Status = "borrowed"
		} else {
			action = "update"
			borrowing.ID = data.ID
			borrowing.Status = "returned"
			borrowing.ReturnedAt = time.Now()
		}

		if action == "create" {
			if err := h.BorrowingRepo.Create(
				[]model.Borrowing{borrowing},
			); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to create borrowing",
				})
			}

			if err := h.BookRepo.Update(
				model.Book{
					ID:     borrowing.BookID,
					Status: "borrowed",
				}); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to update book",
				})
			}
		} else {
			if err := h.BorrowingRepo.Update(
				borrowing,
			); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to update borrowing",
				})
			}

			if err := h.BookRepo.Update(
				model.Book{
					ID:     borrowing.BookID,
					Status: "available",
				}); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to update book",
				})
			}
		}

		return c.JSON(fiber.Map{
			"message":   "Set borrowing successfully",
			"borrowing": borrowing,
		})
	}
}

func (h *BorrowingHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		borrowings, err := h.BorrowingRepo.List()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"borrowings": borrowings,
		})
	}
}

func (h *BorrowingHandler) ListPending() fiber.Handler {
	return func(c *fiber.Ctx) error {
		borrowings, err := h.BorrowingRepo.ListPending()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(nil)
		}

		return c.JSON(fiber.Map{
			"borrowings": borrowings,
		})
	}
}
