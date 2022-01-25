package main

import (
	"gin-blog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterApi(r)
	err := r.Run()
	if err != nil {
		return
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
