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
	hashedPassword := hashPassword(password)
	return &User{
		Firstname: firstname,
		Lastname: lastname,
		Email: email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
} 

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// UUID version 4
	u.ID = uuid.New()
	return
}

func (u *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}