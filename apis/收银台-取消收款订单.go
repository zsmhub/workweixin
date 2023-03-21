package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCloseOrderPaytool 取消收款订单请求
// 文档：https://developer.work.weixin.qq.com/document/path/98046#取消收款订单
type ReqCloseOrderPaytool struct {
	// NonceStr 随机字符串，长度要求在32字节以内，用于保证签名不可预测及防重放攻击。<br/>需保证15分钟内不能重复，推荐随机字符串生成算法
	NonceStr string `json:"nonce_str,omitempty"`
	// OrderID 收款订单号<br/>不多于64字节
	OrderID string `json:"order_id,omitempty"`
	// Sig 数字签名
	Sig string `json:"sig,omitempty"`
	// Ts unix时间戳（中国时区）,精确到秒。<br/>注意业务系统的机器时间与腾讯的时间相差不能超过15分钟
	Ts int `json:"ts,omitempty"`
}

var _ bodyer = ReqCloseOrderPaytool{}

func (x ReqCloseOrderPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCloseOrderPaytool 取消收款订单响应
// 文档：https://developer.work.weixin.qq.com/document/path/98046#取消收款订单
type RespCloseOrderPaytool struct {
	CommonResp
}

var _ bodyer = RespCloseOrderPaytool{}

func (x RespCloseOrderPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCloseOrderPaytool 取消收款订单
// 文档：https://developer.work.weixin.qq.com/document/path/98046#取消收款订单
func (c *ApiClient) ExecCloseOrderPaytool(req ReqCloseOrderPaytool) (RespCloseOrderPaytool, error) {
	var resp RespCloseOrderPaytool
	err := c.executeWXApiPost("/cgi-bin/paytool/close_order", req, &resp, true)
	if err != nil {
		return RespCloseOrderPaytool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCloseOrderPaytool{}, bizErr
	}
	return resp, nil
}
