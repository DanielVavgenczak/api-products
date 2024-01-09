package middleware

import (
	"net/http"
	"strings"

	"github.com/DanielVavgenczak/api-products/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Auth Header
		authHeader := c.GetHeader("Authorization")
		// Must have Bearer in the header
		if !strings.HasPrefix(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is missing",
			})
			return
		}
		// transform AuthHeader in slice
		token := strings.Split(authHeader, " ")
		// Validate, must token have 2 length
		if len(token) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is missing",
			})
			return
		}
		// Token[1] have token acess 
		// Dont must it empty
		if token[1] == " " || token[1] == "undefined"{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is missing",
			})
			return
		}

		// valid token 
		err := helper.ValidToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		
	}
}
