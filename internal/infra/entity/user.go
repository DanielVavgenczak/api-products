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
	
	return &User{
		Firstname: firstname,
		Lastname: lastname,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
} 

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// UUID version 4
	u.ID = uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashPassword)
	return
}

func (u *User) ValidaPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}