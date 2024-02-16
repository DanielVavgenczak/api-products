package repository

import (
	"errors"
	"fmt"

	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"gorm.io/gorm"
)

type CategoryInterface interface {
	/** Title is a new category **/
	Create(title, user_id string)(*entity.Category, error)
	List() ([]entity.Category, error)
	FindByTitleAndUserID(title, user_id string)(*entity.Category, error)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(gorm *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: gorm,
	}
}

func (cateRepo *CategoryRepository) Create(title,user_id string) (*entity.Category, error) {
	newCategory  := entity.NewCategory(title,user_id) 
	var category *entity.Category
	err := cateRepo.DB.Create(newCategory).Scan(&category).Error 
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cateRepo *CategoryRepository) List() ([]entity.Category, error) {
	var categories []entity.Category
	err := cateRepo.DB.Find(&categories).Error
	if err != nil {
		return []entity.Category{}, err
	}
	return categories, nil
}

func (cateRepo *CategoryRepository) FindByTitleAndUserID(title,user_id string) (*entity.Category, error) {
	var category *entity.Category
	err := cateRepo.DB.Find(&category).Where("title AND user_id", title, user_id).Error
	if err != nil {
		return nil, errors.New("erro nessa cara")
	}
	fmt.Println("No repository da categoria", category)
	return category, nil
}


