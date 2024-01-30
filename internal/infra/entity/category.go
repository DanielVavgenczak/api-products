package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(title string) *Category {
	uuidCategory := uuid.New()
	return &Category{
		ID: uuidCategory,
		Title: title,
	}
}
