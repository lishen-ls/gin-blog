package controllers

import "gin-blog/core/controller"
import "gin-blog/core/response"

type IndexController struct {
	controller.Controller
	controller.Interface
}

// HelloWorld
// 输出你好世界
func (c *IndexController) HelloWorld() response.Response {
	return response.Response{
		Code: 200,
		Msg:  "Success",
	}
}
