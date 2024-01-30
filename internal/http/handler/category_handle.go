package handler

import (
	"net/http"

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
	// vamos tentar pegar o context
	c.JSON(http.StatusAccepted, gin.H{
		"data":"ok",
	})
	return
}