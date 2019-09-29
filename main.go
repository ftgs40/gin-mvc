package main

import (
	"github.com/ftgs40/gin-mvc/configs"
	"github.com/ftgs40/gin-mvc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	c := configs.GetServerConfig()
	routes.RegisterRoutes(ginEngine)
	ginEngine.Run(c["PORT"])
}
