package model

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text;" json:"description"`
	Author      string    `gorm:"type:varchar(255)" json:"author"`
	Status      string    `gorm:"type:varchar(100);default:'available'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
