package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetOrderLicense 获取订单详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/95648#获取订单详情
type ReqGetOrderLicense struct {
	// OrderID 订单id，必填
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqGetOrderLicense{}

func (x ReqGetOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetOrderLicense 获取订单详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/95648#获取订单详情
type RespGetOrderLicense struct {
	CommonResp
	Order struct {
		AccountCount struct {
			// BaseCount 基础帐号个数
			BaseCount int `json:"base_count"`
			// ExternalContactCount 互通帐号个数
			ExternalContactCount int `json:"external_contact_count"`
		} `json:"account_count"` // 订单的帐号数详情
		AccountDuration struct {
			// Months 购买的月数，每个月按照31天计算
			Months int `json:"months"`
		} `json:"account_duration"` // 帐号购买时长
		// Corpid 客户企业id，返回加密的corpid
		Corpid string `json:"corpid"`
		// CreateTime 创建时间
		CreateTime int `json:"create_time"`
		// OrderID 订单号
		OrderID string `json:"order_id"`
		// OrderStatus 订单状态，0:待支付，1:已支付，2:未支付，订单已关闭，3:未支付，订单已过期，4:申请退款中，5:退款成功，6:退款被拒绝
		OrderStatus int `json:"order_status"`
		// OrderType 订单类型，1:购买帐号，2:续期帐号 5:历史企业迁移订单
		OrderType int `json:"order_type"`
		// PayTime 支付时间。迁移订单不返回该字段
		PayTime int `json:"pay_time"`
		// Price 订单金额，单位分
		Price int `json:"price"`
	} `json:"order"` // 订单详情
}

var _ bodyer = RespGetOrderLicense{}

func (x RespGetOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetOrderLicense 获取订单详情
// 文档：https://developer.work.weixin.qq.com/document/path/95648#获取订单详情
func (c *ApiClient) ExecGetOrderLicense(req ReqGetOrderLicense) (RespGetOrderLicense, error) {
	var resp RespGetOrderLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_order", req, &resp, true)
	if err != nil {
		return RespGetOrderLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetOrderLicense{}, bizErr
	}
	return resp, nil
}
