package conf

import (
	"flag"
	"fmt"
)

type ENV struct {
	Host string
	AlipayUrl string
	WechatPayQR string
	Port string
	Mode string
}

var Dev =ENV{
	"http://192.168.31.156:2666",
	"HTTPS://QR.ALIPAY.COM/FKX09812XVMA33KQ3IDJ5B",
	"statics/qr.png",
	":2666",
	"debug",

}

var Pro =ENV{
	"https://hipoor.com",
	"HTTPS://QR.ALIPAY.COM/FKX09812XVMA33KQ3IDJ5B",
	"statics/qr.png",
	":2666",
	"release",

}

var CurrentEnv ENV


func init() {
	env:=flag.String("env","Dev","当前的运行环境")
	flag.Parse()
	fmt.Println("当前运行环境",*env)
	if *env=="Dev" {
		CurrentEnv=Dev
	}else {
		CurrentEnv=Pro
	}
}

