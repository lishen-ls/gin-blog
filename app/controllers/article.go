package controllers

import (
	"gin-blog/core/controller"
	"gin-blog/core/response"
)

type ArticleController struct {
	controller.Controller
}

func (c *ArticleController) Get() response.Response {

	return response.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func (c *ArticleController) Post() response.Response {
	return response.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}
