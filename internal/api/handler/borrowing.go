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

func (h *BorrowingHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data types.CreateBorrowingRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		if data.BookID == 0 || data.UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Book and User ID are required",
			})
		}

		var borrowing model.Borrowing

		book, err := h.BookRepo.GetByID(data.BookID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Book not found",
			})
		}
		if book.Status != "available" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Book is not available",
			})
		}

		borrowing.BookID = book.ID

		user, err := h.UserRepo.GetUserById(data.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		if user.Role != "user" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User is not a user",
			})
		}

		borrowing.UserID = data.UserID
		borrowing.Status = "borrowed"

		if err := h.BorrowingRepo.Create(&borrowing); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to create borrowing",
			})
		}

		if err := h.BookRepo.Update(model.Book{
			ID:     book.ID,
			Status: "borrowed",
		}); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update book",
			})
		}

		return c.JSON(fiber.Map{
			"message":   "Borrowing created successfully",
			"borrowing": borrowing,
		})
	}
}

func (h *BorrowingHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data types.UpdateBorrowingRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid body",
			})
		}

		if data.BorrowingID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "id is required",
			})
		}

		var borrowing model.Borrowing

		borrowing.ID = data.BorrowingID

		if data.Status != "" {
			if data.Status != "returned" && data.Status != "borrowed" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Invalid status",
				})
			}
		} else {
			data.Status = "returned"
		}

		borrowing.Status = data.Status
		returnedAt := time.Now()
		borrowing.ReturnedAt = &returnedAt

		if err := h.BorrowingRepo.Update(&borrowing); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update borrowing",
			})
		}

		// if err := h.BookRepo.Update(model.Book{
		// 	ID:     book.ID,
		// 	Status: "borrowed",
		// }); err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"message": "Failed to update book",
		// 	})
		// }

		return c.JSON(fiber.Map{
			"message":   "Borrowing updated successfully",
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
