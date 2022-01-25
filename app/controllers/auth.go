package controllers

import "gin-blog/core/controller"
import "gin-blog/core/response"

type AuthController struct {
	controller.Controller
}

func (c *AuthController) Login() response.Response {
	return response.Response{
		Code: 200,
		Msg:  "登录成功",
	}
}
