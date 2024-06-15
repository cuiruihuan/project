package common

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"src/2110A-zg6/distri/global"
)

var (
	appid      = "9021000131618845"
	privatekey = "MIIEowIBAAKCAQEAoVP+iVhTkzWTQgyy4Xvj635N+/76aTNVJB6cHNrGOL0PevIUX32k/YjINPTkJmoMegfhgg4Bx7KCcHXbNnUUKtu/wthF9vQMYXZItROWuBa9t3fWpOgIny5Br1a8bIs2rhjGnO4sZVl/xSiSv4XAuFvDBLvMnrnBxmv4429oiLlwXSra0SplVf7krLZWhfxwPSgdOhGgz3MWs+9MhuxAq0ui5Jx1fwrzFxErUauTlPZOM78ODppq8OLSCgtKjDt+8jBOT5DqngEA9vEbpDsMzDJ0fFjeYRfOwV/LeAm7Z0Le4Or+cyEZw5P4leEW0qotZwgSxUezhGB1iYJ1ju/q/wIDAQABAoIBABgCwk16weWVAbfA4BbDM6bnA4BP8MEbjL51d9KrlPj42ob35KRZtOavJu0KzbIR8o1vAtcycn1gSh2mzNuMDVZbomzDd5ZLuFQBIM0ifzoi2FWoFmAw0KewYMJR7HVTOu4qbdUJdb4X1fZoFX8/zvEhZyMbOFcfWoHHrCFdgXdntXKrwDAF6vHrHLEkw97HTLJPiearRnS0L2/pnb0zaWbeq9hugeA9okm8YbLnZJV3JIc+2bPWlWgowIomCHOs9Q8lwMQxCPBOIRSCzr0g2lqxmGW5rm4pZQkCObSEXQxmt2Akzm1p7BGd4U0kI2skD7bAUSJkEKooKf0d63jCKRECgYEA2tvvGJ8q+NUrFuTFbytPvXxUmOmTZk44wm13+WAjZYwhYkrH/rO0TiM5Cj3v/aeo+oyi8zHeTwigXZSAowuFdS4C7HSjY75I+GKVK9C93H3snJtWQuPmt1dw2HkCAb5rjSItsSfqA4coNDupXb2w8gTFAVXM3zj94fI07C7OzAMCgYEAvLSyKAjKETmlzjG5qAJagrviyUhTeJo/1ksppZgH1GU9LtR0FPu/j4AH6T6/OJMug7l/f6Cf8OW8X7OSP/HzRrS5OKam2ArxlIEmd8e93x0W/g7VQDX3NvNSjxX3kb7/o3ORMbroT58APaIKWxbFnFBclPXGBYaFX9p+ftwEulUCgYEAgN+nUClRxYrIv8dglxC6+MpAinZoLIL9G0gYsIiy7zXv5pBAajPphWVqTiSgsA3NDBRKU8hWrtVWkSREBz5ejNzWLeSU3cQL9e8fBdj6I98muCUkS2klz2o28dF3pA08CbRJkZr+SquIuhEzrxZnHmw+kYWziPJnVWgpT7ZUYX8CgYAGir0iZiubLURyQYcOLAa3QYw7Lia2p6JO64wmEX8z/c6BMRxVHeUmNGoRy/sVGwUhy+x/oDHrSAufxbnsPZcRfHJvZFekZby6ST/sJyCN+5QwepMqBrTrUsQp8bkDdHuoTvlZoAEtwXVgrZ1uSvvOz3pmL58N0fwNtqVsAHSpgQKBgGkYXcnpvIII3Rq1Bd7ERrhaRHoJ9O0cmdpvu/wq1GGsb7gXmicXzi93nwwuwZitnsJeQYSiam+MblhR57s8sKpovkh8m33p0BaEjiOHBX4/8ue7Rx26iGSg8Fixx7lYW7zXVudb2zhqqviOuhN8jHe3Eq0LxR9JAGhQ2Rht+GxH"
	publickey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwc5gcFN70MFA51/LA+BzNEBU0lh63twu9q1bzikIrrgsvh0G9tMgpf3D3FhiMIaOutaUDdwhI0QMbtEXfD2Rgqfi+ikgVpi9llwLureVl5qXdEX9AWaNbZ8uTqL/WRI3iiY+P8N7l5Zc0o7oqyAiI5NkG2XqbZ0mrwOOC9wOS8Mlmbz+bjSzRFdv5p49Zo9mLZer/oYAgFNMG9I921Ywh7q88fn23+2eyYKfOW4hNnTYlHpvji9tUl0rA24ZvC+kouYpMnGAdFHElQU1uTaY2d7FHL4nPmi6e2eMuBi3mOGH3ujSNJHn/StM2voswM239i4ETI2x9a6FPtnoco8/IQIDAQAB"
)

func NewPayClient() *alipay.Client {
	var err error
	if global.AliPayClient, err = alipay.New(appid, privatekey, false); err != nil {
		zap.S().Info("警告::: ")
	}
	//加载公钥
	global.AliPayClient.LoadAliPayPublicKey(publickey)
	return global.AliPayClient
}

// 订单支付
func AliPay(orderNumber string, price string) string {
	NewPayClient()
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://6915c1b6.r18.cpolar.top" + "/order/notify"
	p.ReturnURL = "http://www.baidu.com"
	p.Subject = "支付测试:" + orderNumber
	p.OutTradeNo = orderNumber
	p.TotalAmount = price
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, _ := global.AliPayClient.TradePagePay(p)
	return url.String()
}

func Refund(orderNumber string, price string) (*alipay.TradeRefundRsp, error) {
	// 创建或获取支付宝支付客户端的实例
	payClient := NewPayClient()

	// 构建退款请求
	refundRequest := alipay.TradeRefund{
		TradeNo:      orderNumber,
		RefundAmount: price,
	}

	// 调用退款服务
	result, err := payClient.TradeRefund(refundRequest)
	if err != nil {
		// 记录错误并返回
		zap.S().Errorf("退款失败: OrderNumber: %s, Error: %v", orderNumber, err)
		return nil, err
	}

	// 检查退款结果
	if !result.IsSuccess() {
		// 根据result中的信息记录失败详情
		zap.S().Errorf("退款请求失败: OrderNumber: %s, Response: %+v", orderNumber, result)
		return result, fmt.Errorf("退款失败，详情请查看日志")
	}

	// 返回成功的退款结果
	return result, nil
}
