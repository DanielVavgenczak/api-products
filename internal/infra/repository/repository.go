package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepository *UserRepository
}

// InitRepository should be called in main.go 
func InitRepository(db *gorm.DB) *Repositories {	
	userRepo := NewUserRepository(db)
	return &Repositories{
		UserRepository: userRepo,
	}
}