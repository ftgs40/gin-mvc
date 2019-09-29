package routes

import (
	"github.com/ftgs40/gin-mvc/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes is set routing project function
func RegisterRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.GET("/", controllers.HomeIndex)
	router.GET("/api", controllers.APIExample)
}
