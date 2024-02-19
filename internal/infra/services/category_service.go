package services

import (
	"errors"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
)

var (
	ErrCategoryIsAlreadRegistedWhitThisUser = errors.New("category is alread exists with this user")
)
type CategoryService struct {
	repository repository.CategoryInterface
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: &repo,
	}
}

func (service *CategoryService) CreateCategory(categoryInput dto.CategoryInput)(*entity.Category, error) {
	categoryAlreadExists, _ := service.repository.FindByTitleAndUserID(categoryInput.Title, categoryInput.UserID) 
	if categoryAlreadExists != nil && categoryAlreadExists.Title == categoryInput.Title  {
		return nil, ErrCategoryIsAlreadRegistedWhitThisUser
	}
	category, err := service.repository.Create(categoryInput.Title, categoryInput.UserID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) FindCategoryByUser(user_id string)([]*entity.Category, error) {
	categories, err := s.repository.ListByUser(user_id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}