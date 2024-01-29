package services

import (
	"errors"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/helper"
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
	"github.com/google/uuid"
)

var (
	ErrUserAlreadExists = errors.New("user is alread exists")
	ErrUserNotFound = errors.New("user not found")
)

type UserResponse struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Token string `json:"token_acess"`
} 

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
	if userExists.Email  == userInput.Email && userExists.ID != uuid.Nil {
		return nil, ErrUserAlreadExists
	}
	user, err := service.repository.Create(userInput.Firstname, userInput.Lastname, userInput.Email, userInput.Password, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) FindByEmailUser(email string) (*entity.User, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if len(user.Email) == 0  {
		return nil, err
	}
	return user, nil
}


func (service *UserService) FindByIDUser(id string) (entity.User, error) {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (service *UserService) UserLogin(email, password string) (*UserResponse, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	// compare password 
  if ok := user.ValidPassword(password); !ok {
		return nil, errors.New("user not found")
	}
	token, err := helper.GenerateToken(user.ID.String(), 300)
	if err != nil {
		return nil, err
	}
	userResponse := UserResponse{
		ID: user.ID.String(),
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		Token: token,
	}

	return &userResponse, nil

}