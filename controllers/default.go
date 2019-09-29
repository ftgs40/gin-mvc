package controllers

import (
	"github.com/gin-gonic/gin"
)

// APIExample is api example
func APIExample(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "render text in default",
	})
}
