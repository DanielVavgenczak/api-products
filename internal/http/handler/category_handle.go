package handler

import (
	"net/http"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/infra/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandle struct {
	cateService *services.CategoryService
}

func NewCategoryHandle(cateService services.CategoryService) *CategoryHandle {
	return &CategoryHandle{
		cateService: &cateService,
	}
}

func (cateHandle *CategoryHandle) CreateCategoryHandler(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var categoryInput dto.CategoryInput
	categoryInput.UserID = user_id.(string)
	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}				
	categoryUser, err := cateHandle.cateService.CreateCategory(categoryInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code":http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status_code":http.StatusAccepted,
		"data": categoryUser,
	})
}