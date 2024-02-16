package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID uuid.UUID `json:"id" gorm:"primarykey"`
	Title string `json:"title" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id" gorm:"type:char(36);primary_key"`
}

func NewCategory(title, user_id string) *Category {
	uuidCategory := uuid.New()
	userIdParsed,err := uuid.Parse(user_id)
	if err != nil {
		panic(err.Error())
	}
	return &Category{
		ID: uuidCategory,
		Title: title,
		UserID: userIdParsed,
	}
}

func (c *Category) BeforeSave(tx *gorm.DB) {
	c.ID = uuid.New()
}