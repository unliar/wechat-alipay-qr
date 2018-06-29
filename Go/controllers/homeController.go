package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetPath() string {
	return "abc"
}

func HomeRoute(c *gin.Context) {

	h := c.GetHeader("user-agent")

	isWechat := strings.Contains(h, "MicroMessenger")
	isAlipay := strings.Contains(h, "AlipayClient")
	// 如果是微信就跳转到/wechat渲染
	if isWechat {
		c.Redirect(200, "/wechat")
		return
	}

	if isAlipay {
		c.Redirect(200, "http://hipoor.com")
		return
	}

	if !isAlipay && !isWechat {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"message": "嘻嘻默认首页",
		})
		return
	}

}
