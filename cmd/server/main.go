package main

import (
	"fmt"

	"github.com/DanielVavgenczak/api-products/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()

	fmt.Println(db.DryRun)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() 
}