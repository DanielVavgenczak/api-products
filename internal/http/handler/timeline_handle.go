package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TimelineHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":"timeline",
	})}