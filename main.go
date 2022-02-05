package main

import (
	"gin-blog/bootstrap"
	"gin-blog/global"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()
	// 初始化日志
	bootstrap.InitLog()
	// 初始化数据库
	bootstrap.InitDB()
	// 注册验证器
	bootstrap.InitValidator()
	// 启动服务器
	bootstrap.RunServer()

	global.App.Log.Info("服务开启")
}
