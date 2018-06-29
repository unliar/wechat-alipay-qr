package controllers

import "github.com/gin-gonic/gin"

func WechatRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is from wechat router~",
	})
}
