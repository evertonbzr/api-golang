package model

import (
	"time"
)

type Borrowing struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint       `gorm:"index" json:"user_id"`
	BookID     uint       `gorm:"index" json:"book_id"`
	BorrowedAt time.Time  `gorm:"autoCreateTime" json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"`
	Status     string     `gorm:"type:varchar(100),default:'borrowed'" json:"status"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
