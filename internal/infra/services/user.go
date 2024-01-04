package services

import (
	"errors"
	"fmt"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
)

type UserService struct {
	repository repository.UserInterface
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repository: &repo,
	}
}

func (service *UserService) CreateUser(userInput dto.UserInput) (*entity.User, error){
	userExists, err := service.repository.FindByEmail(userInput.Email)
	if err != nil {
		return nil, err
	}
	if userExists.Email  == userInput.Email && userExists.ID != "" {
		fmt.Println("de bas", userExists.Email)
		return nil, errors.New("user is alread exists")
	}
	user, err := service.repository.Create(userInput.Firstname, userInput.Lastname, userInput.Email, userInput.Password, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) FindByEmailUser(email string) (entity.User, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}