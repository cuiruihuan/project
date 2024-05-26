package common

import (
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"src/2110A-zg6/week1/global"
)

func NewPayClient() *alipay.Client {
	var err error
	if global.AliPayClient, err = alipay.New(global.ConfigAll.Alipay.Appid, global.ConfigAll.Alipay.Privatekey, false); err != nil {
		zap.S().Info("警告::: ")
	}
	//加载公钥
	global.AliPayClient.LoadAliPayPublicKey(global.ConfigAll.Alipay.Publickey)
	return global.AliPayClient
}

// 订单支付
func AliPay(orderName string, price string) string {
	NewPayClient()
	var p = alipay.TradePagePay{}
	p.NotifyURL = "https://1b21610d.r7.cpolar.top" + "/pay/notify"
	p.ReturnURL = "http://www.baidu.com"
	p.Subject = "支付测试:" + orderName
	p.OutTradeNo = orderName
	p.TotalAmount = price
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, _ := global.AliPayClient.TradePagePay(p)
	return url.String()
}
