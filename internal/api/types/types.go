package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Author      string `json:"author"`
}

type CreateBorrowingRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

type UpdateBorrowingRequest struct {
	BorrowingID uint   `json:"id"`
	Status      string `json:"status"`
}
