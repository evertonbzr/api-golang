package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName  string    `gorm:"type:varchar(255)" json:"full_name"`
	Email     string    `gorm:"type:varchar(255);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
