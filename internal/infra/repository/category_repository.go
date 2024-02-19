package repository

import (
	"errors"

	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
	"gorm.io/gorm"
)

type CategoryUserResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

type CategoryInterface interface {
	/** Title is a new category **/
	Create(title, user_id string)(*entity.Category, error)
	ListByUser(user_id string) ([]*entity.Category, error)
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
	newCategory := entity.NewCategory(title,user_id) 
	var category *entity.Category
	err := cateRepo.DB.Create(newCategory).Scan(&category).Error 
	if err != nil || err == gorm.ErrDuplicatedKey {
		return nil, errors.New("this user is alread register this category")
	}
	return category, nil
}

func (cateRepo *CategoryRepository) ListByUser(user_id string) ([]*entity.Category, error) {
	var categories []*entity.Category
	//Joins("User", cateRepo.DB.Where(&entity.User{ID: convert})). // Specify users table
	//convert, _ := uuid.Parse(user_id)
	err := cateRepo.DB.Table("categories").
				Preload("User",func (tx *gorm.DB) *gorm.DB {
					return tx.Select("id","firstname")
				}).
				Where("user_id = ?", user_id).
        Find(&categories).
        Error
    if err != nil {
        return nil, err
    }
    return categories, nil
}

func (cateRepo *CategoryRepository) FindByTitleAndUserID(title,user_id string) (*entity.Category, error) {
	var category *entity.Category
	err := cateRepo.DB.
		Model(&entity.Category{}).
		Where("title = ? AND user_id = ?", title, user_id).
		Find(&category). 
		Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return category, nil
}


