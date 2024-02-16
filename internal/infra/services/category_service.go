package services

import (
	"errors"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
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
	categoryExists, err := service.repository.FindByTitleAndUserID(categoryInput.Title, categoryInput.UserID)
	if err != nil {
		return nil, err
	}

	if categoryExists.Title == categoryInput.Title {

			return nil, errors.New("categoria is aread exists")
	}
	category, err := service.repository.Create(categoryInput.Title, categoryInput.UserID)
	if err != nil {
		return nil, err
	}
	return category, nil
}