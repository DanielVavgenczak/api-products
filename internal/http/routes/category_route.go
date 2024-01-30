package routes

import (
	"github.com/DanielVavgenczak/api-products/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rg *gin.RouterGroup, categoryHandler handler.CategoryHandle) {
	category := rg.Group("/category")

	category.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data":"category",
		})
	})
	return 
}