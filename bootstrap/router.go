package bootstrap

import (
	"gin-blog/global"
	"gin-blog/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api")
	routes.SetApiRoutes(apiGroup)

	return router
}

func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}
