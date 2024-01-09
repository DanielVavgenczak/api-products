package handler

import (
	"net/http"

	"github.com/DanielVavgenczak/api-products/internal/dto"
	"github.com/DanielVavgenczak/api-products/internal/helper"
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

// CreateUser godoc
// @Summary Register User
// @Description create a new user account
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.UserInput true "user request"
// @Success 201 
// @Failure 400
// @Router /api/v1/user/ [post]
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

// Find User By ID godoc
// @Summary Find user by id
// @Description Find user by id
// @Tags user
// @Accept json
// @Produce json
// @Param  id path  string  false "User ID" Format(uuid)
// @Success 200 
// @Failure 404
// @Router /api/v1/user/{id} [get]
func (handler *UserHandler) HandleFindByID(c *gin.Context) {
	id := c.Param("id")
	user,err := handler.userService.FindByIDUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

	// Login godoc
	// @Summary Login User account
	// @Description Login user account
	// @Tags user
	// @Accept json
	// @Produce json
	// @Param request body dto.UserInputLogin true "user request"
	// @Success 201 
	// @Failure 401
	// @Router /api/v1/login [post]
	func (handler *UserHandler) HandleLogin(c *gin.Context) {
		var login dto.UserInputLogin 
		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}		
		user, err := handler.userService.FindByEmailUser(login.Email)	
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// compare password 
		if ok := user.ValidaPassword(login.Password); !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user not found",
			})
			return
		}

		token, err := helper.GenerateToken(user.ID.String(), 300)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"data": token,
		})
	}

	