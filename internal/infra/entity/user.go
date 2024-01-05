package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"primarykey"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email" gorm:"unique"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(firstname,lastname, email, avatar, password string ) *User{
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		panic(err)
	}
	return &User{
		Firstname: firstname,
		Lastname: lastname,
		Email: email,
		Password: string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
} 

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u.ID = uuid.New()
	return
}