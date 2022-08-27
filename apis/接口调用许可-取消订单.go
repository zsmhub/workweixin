package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCancelOrderLicense 取消订单请求
// 文档：https://developer.work.weixin.qq.com/document/path/96106#取消订单
type ReqCancelOrderLicense struct {
	Corpid  string `json:"corpid"`
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqCancelOrderLicense{}

func (x ReqCancelOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCancelOrderLicense 取消订单响应
// 文档：https://developer.work.weixin.qq.com/document/path/96106#取消订单
type RespCancelOrderLicense struct {
	CommonResp
}

var _ bodyer = RespCancelOrderLicense{}

func (x RespCancelOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCancelOrderLicense 取消订单
// 文档：https://developer.work.weixin.qq.com/document/path/96106#取消订单
func (c *ApiClient) ExecCancelOrderLicense(req ReqCancelOrderLicense) (RespCancelOrderLicense, error) {
	var resp RespCancelOrderLicense
	err := c.executeWXApiPost("/cgi-bin/license/cancel_order", req, &resp, true)
	if err != nil {
		return RespCancelOrderLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCancelOrderLicense{}, bizErr
	}
	return resp, nil
}
