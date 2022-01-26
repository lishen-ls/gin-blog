package main

import (
	"gin-blog/bootstrap"
	"gin-blog/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()
	// 初始化日志
	bootstrap.InitLog()
	// 初始化数据库
	bootstrap.InitDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		db, _ := global.App.DB.DB()
		row := db.QueryRow("SELECT * FROM user WHERE id = ?", "4494fa8c-580f-4781-97aa-7ee253b9a008")
		row.Scan()
		c.String(http.StatusOK, "pong")
	})

	global.App.Log.Info("服务开启")

	r.Run(":" + global.App.Config.App.Port)
	//routes.RegisterApi(r)
	//err := r.Run()
	//if err != nil {
	//	return
	//}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
