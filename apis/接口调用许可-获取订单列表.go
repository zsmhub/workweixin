package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListOrderLicense 获取订单列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/95647#获取订单列表
type ReqListOrderLicense struct {
	// Corpid 企业id，只支持加密的corpid。若指定corpid且corpid为服务商测试企业，则返回的订单列表为测试订单列表。否则只返回正式订单列表
	Corpid string `json:"corpid,omitempty"`
	// StartTime 开始时间,下单时间。可不填。但是不能单独指定该字段，start_time跟end_time必须同时指定。
	StartTime int `json:"start_time,omitempty"`
	// EndTime 结束时间,下单时间。起始时间跟结束时间不能超过31天。可不填。但是不能单独指定该字段，start_time跟end_time必须同时指定。
	EndTime int `json:"end_time,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500
	Limit int `json:"limit,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
}

var _ bodyer = ReqListOrderLicense{}

func (x ReqListOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListOrderLicense 获取订单列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/95647#获取订单列表
type RespListOrderLicense struct {
	CommonResp
	// HasMore 是否结束
	HasMore int `json:"has_more"`
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
	OrderList  []struct {
		// OrderID 订单id
		OrderID string `json:"order_id"`
		// OrderType 订单类型，1:购买帐号，2:续期帐号，5:历史企业迁移订单
		OrderType int `json:"order_type"`
	} `json:"order_list"` // 订单列表
}

var _ bodyer = RespListOrderLicense{}

func (x RespListOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListOrderLicense 获取订单列表
// 文档：https://developer.work.weixin.qq.com/document/path/95647#获取订单列表
func (c *ApiClient) ExecListOrderLicense(req ReqListOrderLicense) (RespListOrderLicense, error) {
	var resp RespListOrderLicense
	err := c.executeWXApiPost("/cgi-bin/license/list_order", req, &resp, true)
	if err != nil {
		return RespListOrderLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListOrderLicense{}, bizErr
	}
	return resp, nil
}
