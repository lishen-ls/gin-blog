package controller

import "github.com/gin-gonic/gin"

import "gin-blog/core/response"

type Controller struct {
	Prefix string
	Interface
}

type Action func(ctx *gin.Context) response.Response

type Interface interface {
	Get(c *gin.Context) response.Response
	Post(c *gin.Context) response.Response
	Delete(c *gin.Context) response.Response
	Put(c *gin.Context) response.Response
}
