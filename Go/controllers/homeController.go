package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wechat-alipay-qr/Go/conf"
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
		c.Redirect(301, "/wechat")
		return
	}

	if isAlipay {
		c.Redirect(301, conf.CurrentEnv.Host)
		return
	}

	if !isAlipay && !isWechat {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"message": "嘻嘻默认首页",
		})
		return
	}

}
