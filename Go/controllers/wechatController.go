package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-alipay-qr/Go/conf"
)

func WechatRoute(c *gin.Context) {
	c.HTML(http.StatusOK,"wechat.tmpl",gin.H{
		"url":conf.CurrentEnv.WechatPayQR,
	})
}
