package repository

import (
	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"gorm.io/gorm"
)

type CategoryInterface interface {
	/** Title is a new category **/
	Create(title string)(*entity.Category, error)
	List() ([]entity.Category, error)
	FindByTitle(idadmin, title string)(*entity.Category, error)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(gorm *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: gorm,
	}
}

func (cateRepo *CategoryRepository) Create(title string) (*entity.Category, error) {
	newCategory  := entity.NewCategory(title) 
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

func (cateRepo *CategoryRepository) FindByTitle(idAdmin, title string) (*entity.Category, error) {
	var categories entity.Category
	err := cateRepo.DB.Find(&categories).Where("admin_id AND title", idAdmin, title).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}


