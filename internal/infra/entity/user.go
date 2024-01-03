package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string `json:"id" gorm:"primarykey"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
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
		ID: uuid.NewString(),
		Firstname: firstname,
		Lastname: lastname,
		Email: email,
		Password: string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
} 