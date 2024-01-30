package services

import (
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
)

type Services struct {
	Repositories repository.Repositories
	UserService  UserService
	CategoryService CategoryService
}

func InitServices(repos repository.Repositories) *Services{
	userService := NewUserService(*repos.UserRepository)
	categoryService := NewCategoryService(*repos.CategoryRepository)
	return &Services{
		Repositories: repos,
		UserService: *userService,
		CategoryService: *categoryService,
	}
}
