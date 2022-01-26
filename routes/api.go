package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(group gin.RouterGroup) {
	group.GET("/ping", func(ct *gin.Context) {
		ct.String(http.StatusOK, "test")
	})
}
