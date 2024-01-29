package repository

import (
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(firstname, lastname, email, password, avatar string) (*entity.User, error)
	FindByEmail(email string) (*entity.User,error)
	FindByID(email string) (entity.User,error)
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

func (repo *UserRepository) FindByEmail(email string)(*entity.User,error){
	var user entity.User
	err := repo.db.Table("users").Where("email", email).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (repo *UserRepository) FindByID(id string)(entity.User,error){
	var user entity.User
	convertID, err := uuid.Parse(id)
	if err != nil {
		return entity.User{}, err
	}
	err = repo.db.Raw("SELECT * FROM users WHERE id = ? ", convertID).Scan(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}