package main

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
	"wechat-qr-Go/app/router"
)

func main() {

	err := qr.WriteFile("http://hipoor.com", qr.Medium, 256, "statics/qrImg.png")
	fmt.Println(err)
	r := router.RouterStates()
	r.LoadHTMLGlob("./views/*")
	r.Run()
}
