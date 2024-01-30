package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepository *UserRepository
	CategoryRepository *CategoryRepository
}

// InitRepository should be called in main.go 
func InitRepository(db *gorm.DB) *Repositories {	
	userRepo := NewUserRepository(db)
	cateRepo := NewCategoryRepository(db)
	return &Repositories{
		UserRepository: userRepo,
		CategoryRepository: cateRepo,
	}
}