package services

import (
	"context"

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

func (service *CategoryService) CreateCategory(ctx context.Context, title string)(*entity.Category, error) {
		
	//exists, err : service.repository.FindByTitle()
	return nil, nil
}