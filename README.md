# wechat-alipay-qr
use one qr picture to guide to your wechat or alipay account.

微信支付宝二维码聚合,识别不同的user-agent重定向到支付界面

## Node.js版本

### 启动 

```
cd Node.js/

npm i

npm start

或者全局安装nodemon

npm i nodemon -g

npm i

npm start:hot

```

### 开发
> 编辑 Node.js/config/index.js


## Go语言版本
>web框架 [gin](https://github.com/gin-gonic/gin)

### 启动
```
cd Go/

// 依赖安装
govendor sync

// 启动本地测试环境

go run app.go -env=Dev

// 启动生产环境

go run app.go -env=Pro 

```

### 部署

```
// 编译
go build app.go

// 记得复制views文,statics文件夹到执行文件同一目录

// 测试环境部署
./app -env=Dev

// 正式环境部署
./app -env=Pro

```




