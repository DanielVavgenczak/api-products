package main

import (
	"log"

	"github.com/DanielVavgenczak/api-products/docs"
	"github.com/DanielVavgenczak/api-products/internal/config"
	"github.com/DanielVavgenczak/api-products/internal/http/handler"
	"github.com/DanielVavgenczak/api-products/internal/infra/database"
	"github.com/DanielVavgenczak/api-products/internal/infra/repository"
	"github.com/DanielVavgenczak/api-products/internal/infra/services"
	"github.com/DanielVavgenczak/api-products/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Api Products
// @version 1.0
// @description Creating simple api product.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1
// @schemes http
func main() {
	r := gin.Default()
	ok := config.LoadEnv()
	if !ok {
		log.Fatal("error in load env")
	}
	db := database.InitDB()
	repositories := repository.InitRepository(db)
	services := services.InitServices(*repositories)
	
	userHandler := handler.NewUserHandler(services.UserService)

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")	
	v1.POST("/login", userHandler.HandleLogin)
	v1.POST("/user", userHandler.HandlerCreateUser)
	v1.GET("/user/:id", middleware.Authentication(), userHandler.HandleFindByID)
	// badgg.irlk
	urlDoc := ginSwagger.URL("http://localhost:8080/docs/doc.json") 
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,urlDoc))
	r.Run() 
}