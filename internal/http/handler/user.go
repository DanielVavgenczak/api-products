package handler

import (
	"net/http"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/infra/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler{
	return &UserHandler{
		userService: &service,
	}
}

func (handler *UserHandler) HandlerCreateUser(c *gin.Context) {
	var userInput dto.UserInput 
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	user,err := handler.userService.CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": user,
	})

}
