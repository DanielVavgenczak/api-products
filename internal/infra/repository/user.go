package repository

import (
	"fmt"

	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(firstname, lastname, email, password, avatar string) (*entity.User, error)
	FindByEmail(email string) (entity.User,error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(firstname, lastname, email, password, avatar string) (*entity.User, error) {
	newUser := entity.NewUser(
		firstname,
		lastname,
		email,
		avatar,
		password,
	)
	var user entity.User
	err := repo.db.Create(&newUser).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string)(entity.User,error){
	var user entity.User
	err := repo.db.Raw("SELECT * FROM users WHERE email = ? ", email).Scan(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	fmt.Println("vazio?  ", user)
	return user, nil
}