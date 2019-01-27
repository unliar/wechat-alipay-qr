<template>
  <div class="qr-code">
    <h1>你看到我掉的一百万了吗</h1>
    <h2 v-show="type == 'wechat'">不长按一下，你不知道你有多穷！</h2>
    <h2 v-show="type == 'alipay'">给尊贵的支付宝用户递可乐！</h2>
    <qr :value="showQRSource" :options="{ width: 560 }" tag="img" v-show="type != 'alipay'"></qr>
  </div>
</template>
<script>
import qr from "@chenfengyuan/vue-qrcode";

export default {
  data() {
    return {
      wechat: "wxp://f2f0Ow0E41-fNXXPOGN5Pu4J5E02jnhgTTwh",
      alipay: "https://qr.alipay.com/tsx01801kzs9lew0c8x0484",
      qq:
        "https://i.qianbao.qq.com/wallet/sqrcode.htm?m=tenpay&f=wallet&a=1&ac=B9ADDB2B87E69E7BF2720F6E05A2AC13BD56123CEEF7FED9F62A23CA0F707147&u=370732889&n=%E8%BF%9C%E6%B5%85",
      showQRSource: "https://hipoor.com",
      type: ""
    };
  },
  methods: {
    GerCurrentEnv() {
      if (window && window.navigator && window.navigator.userAgent) {
        const isWechat = window.navigator.userAgent.includes("MicroMessenger");
        const isAlipay = window.navigator.userAgent.includes("AlipayClient");
        const isQQ = window.navigator.userAgent.includes("MQQBrowser");
        if (isWechat) {
          this.showQRSource = this.wechat;
          this.type = "wechat";
          return;
        }
        if (isQQ) {
          this.showQRSource = this.qq;
          this.type = "wechat";
          return;
        }
        if (isAlipay) {
          this.type = "alipay";
          setTimeout(() => {
            window.open(this.alipay);
          }, 1000);
          return;
        }
        this.showQRSource = `${window.location.href}`;
      }
      this.showQRSource = `${window.location.href}`;
    }
  },
  mounted() {
    this.GerCurrentEnv();
  },
  components: {
    qr
  }
};
</script>



<style scoped>
* {
  margin: 0;
  padding: 0;
}
body {
  overflow: hidden;
  text-align: center;
  font-size: 1rem;
  max-width: 100%;
}
.qr-code h1,
.qr-code h2 {
  text-align: center;
  font-weight: normal;
  font-size: 3rem;
  padding: 15px;
}
.qr-code img {
  display: block;
  margin: 15px auto;
}
</style>
