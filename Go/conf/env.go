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
	"127.0.0.1",
	"HTTPS://QR.ALIPAY.COM/FKX0736735EJNZE1PEYUA6",
	"statics/qr.png",
	":2666",
	"debug",

}

var Pro =ENV{
	"https://hipoor.com",
	"HTTPS://QR.ALIPAY.COM/FKX0736735EJNZE1PEYUA6",
	"statics/qr.png",
	":2666",
	"release",

}

var CurrentEnv ENV

var flagEnv string

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

