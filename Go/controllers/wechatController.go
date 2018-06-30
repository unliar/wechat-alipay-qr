package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WechatRoute(c *gin.Context) {
	c.HTML(http.StatusOK,"wechat.tmml",gin.H{
		"url":"/statics/qr.png",
	})
}
