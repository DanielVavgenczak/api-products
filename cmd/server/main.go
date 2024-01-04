package main

import (
	"github.com/DanielVavgenczak/api-products/internal/http/handler"
	"github.com/DanielVavgenczak/api-products/internal/infra/database"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
	"github.com/DanielVavgenczak/api-products/internal/infra/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()

	repositories := repository.InitRepository(db)
	services := services.InitServices(*repositories)
	
	userHandler := handler.NewUserHandler(services.UserService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user", userHandler.HandlerCreateUser)
	r.Run() 
}