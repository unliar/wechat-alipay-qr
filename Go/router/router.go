package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat-qr-Go/app/controllers"
)

func RouterStates() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/statics", http.Dir("./statics"))
	router.GET("/", controllers.HomeRoute)
	router.GET("/wechat", controllers.WechatRoute)
	return router
}
