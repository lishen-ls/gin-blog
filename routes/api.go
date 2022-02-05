package routes

import (
	"gin-blog/app/common/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiRoutes(group *gin.RouterGroup) {
	group.GET("/ping", func(ct *gin.Context) {
		ct.String(http.StatusOK, "test")
	})

	group.POST("/article", func(c *gin.Context) {
		var form request.Article

		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": request.GetErrorMsg(form, err)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
}
