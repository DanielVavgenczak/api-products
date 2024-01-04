package services

import (
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
)

type Services struct {
	Repositories repository.Repositories
	UserService  UserService
}

func InitServices(repos repository.Repositories) *Services{
	userService := NewUserService(*repos.UserRepository)
	return &Services{
		Repositories: repos,
		UserService: *userService,
	}
}
