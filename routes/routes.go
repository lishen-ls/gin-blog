package routes

import (
	"github.com/gin-gonic/gin"
)
import (
	"gin-blog/app/controllers"
)

func index(r *gin.Engine) {
	indexController := &controllers.IndexController{}

	r.GET("/:id", func(c *gin.Context) {
		response := indexController.HelloWorld()
		c.JSON(200, response)
	})
}

func auth(r *gin.Engine) {
	authController := &controllers.AuthController{}

	r.POST("/auth/login", func(c *gin.Context) {
		response := authController.Login()
		c.JSON(200, response)
	})
}

func article(r *gin.Engine) {
	articleController := &controllers.ArticleController{}
	r.GET("/article/:id", func(ctx *gin.Context) {
		response := articleController.Get()
		ctx.JSON(response.Code, response)
	})
	r.GET("/article", func(ctx *gin.Context) {
		response := articleController.Get()
		ctx.JSON(response.Code, response)
	})
	r.DELETE("/article/:id", func(ctx *gin.Context) {
		response := articleController.Delete(ctx)
		ctx.JSON(response.Code, response)
	})
	r.POST("/article/", func(ctx *gin.Context) {
		response := articleController.Post()
		ctx.JSON(response.Code, response)
	})
	r.PUT("/article/:id", func(ctx *gin.Context) {
		response := articleController.Put(ctx)
		ctx.JSON(response.Code, response)
	})
}

func RegisterApi(r *gin.Engine) {
	// 添加index路由
	index(r)
	// 添加auth路由
	auth(r)
	// 添加article路由
	article(r)
}
