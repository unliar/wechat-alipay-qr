let express = require('express');
let router = express.Router();
let qr = require("qr-image")
/* 配置你使用的可访问地址. */

let qrImg = qr.image("http://127.0.0.1:2666", {
  type: 'png'
})
qrImg.pipe(require("fs").createWriteStream("public/img/qrImg.png"))
/* 首页 */
router.get('/', function (req, res) {
  const {
    'user-agent': agent
  } = req.headers

  if (agent.indexOf("MicroMessenger") !== -1) {
    //跳转微信
    res.redirect("/wechat")
  } else if (agent.indexOf("AlipayClient") !== -1) {
    //支付宝扫描直接跳转到转账
    res.redirect('HTTPS://QR.ALIPAY.COM/FKX0736735EJNZE1PEYUA6') //你的支付宝收款识别出的地址
  } else {
    res.render("index")
  }

});

router.get("/wechat", (req, res) => {
  res.render('wechat', {
    url: 'img/qr.png' //你的微信转账图片地址
  })
})
module.exports = router;