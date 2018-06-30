package main

import (
	qr "github.com/skip2/go-qrcode"
	"wechat-alipay-qr/Go/router"
	"wechat-alipay-qr/Go/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(conf.CurrentEnv.Mode)
	// 生成首页的图片
	qr.WriteFile(conf.CurrentEnv.Host, qr.Medium, 256, "statics/qrImg.png")
	r := router.RouterStates()
	r.LoadHTMLGlob("views/*")
	r.Run(conf.CurrentEnv.Port)
}
