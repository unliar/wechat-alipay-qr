let express = require('express');
let router = express.Router();
let qr = require("qr-image")
let h = {
  hostName,
  alipayUrl,
  wechatPayQR
} = require("../config/index")


let qrImg = qr.image(hostName, {
  type: 'png'
})
qrImg.pipe(require("fs").createWriteStream("public/img/qrImg.png"))

router.get('/', function (req, res) {
  const {
    'user-agent': agent
  } = req.headers

  if (agent.indexOf("MicroMessenger") !== -1) {
    //跳转微信
    res.redirect("/wechat")
  } else if (agent.indexOf("AlipayClient") !== -1) {
    //支付宝扫描直接跳转到转账
    res.redirect(alipayUrl)
  } else {
    //否则渲染入口页
    res.render("index")
  }

});

router.get("/wechat", (req, res) => {
  res.render('wechat', {
    url: wechatPayQR //你的微信转账图片地址
  })
})
module.exports = router;