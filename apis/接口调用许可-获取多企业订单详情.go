package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetUnionOrderLicense 获取多企业订单详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/98893#获取多企业订单详情
type ReqGetUnionOrderLicense struct {
	// Cursor 用于分页查询的游标，字符串类型，填写上一次调用返回的 next_cursor，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500
	Limit int `json:"limit,omitempty"`
	// OrderID 订单id，必填
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqGetUnionOrderLicense{}

func (x ReqGetUnionOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetUnionOrderLicense 获取多企业订单详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/98893#获取多企业订单详情
type RespGetUnionOrderLicense struct {
	BuyList []struct {
		AccountCount struct {
			// BaseCount 基础帐号个数
			BaseCount int `json:"base_count"`
			// ExternalContactCount 互通帐号个数
			ExternalContactCount int `json:"external_contact_count"`
		} `json:"account_count"` // 订单的帐号数详情
		AccountDuration struct {
			// Days 购买的天数
			Days int `json:"days"`
			// Months 购买的月数，每个月按照31天计算
			Months int `json:"months"`
		} `json:"account_duration"` // 帐号购买时长
		// Corpid 客户企业id，返回加密的corpid
		Corpid string `json:"corpid"`
		// SubOrderID 子订单id，可以调用<a href="#38201" rel="nofollow">获取订单中的帐号列表接口</a>以获取帐号列表
		SubOrderID string `json:"sub_order_id"`
	} `json:"buy_list"` // 多企业购买信息列表
	CommonResp
	// HasMore 是否有更多。 0: 没有， 1: 有
	HasMore int `json:"has_more"`
	// NextCursor 分页游标，下次请求时填写到 cursor 以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
	Order      struct {
		// CreateTime 创建时间，Unix时间戳
		CreateTime int `json:"create_time"`
		// OrderID 订单号
		OrderID string `json:"order_id"`
		// OrderStatus 订单状态<br/>0:待支付<br/>1:已支付<br/>2:已取消（未支付，订单已关闭）<br/>3:未支付，订单已过期
		OrderStatus int `json:"order_status"`
		// OrderType 订单类型<br/>8:多企业新购订单
		OrderType int `json:"order_type"`
		// PayTime 支付时间，Unix时间戳
		PayTime int `json:"pay_time"`
		// Price 订单金额，单位分
		Price int `json:"price"`
	} `json:"order"` // 订单详情
}

var _ bodyer = RespGetUnionOrderLicense{}

func (x RespGetUnionOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUnionOrderLicense 获取多企业订单详情
// 文档：https://developer.work.weixin.qq.com/document/path/98893#获取多企业订单详情
func (c *ApiClient) ExecGetUnionOrderLicense(req ReqGetUnionOrderLicense) (RespGetUnionOrderLicense, error) {
	var resp RespGetUnionOrderLicense
	err := c.executeWXApiPost("/cgi-bin/license/get_union_order", req, &resp, true)
	if err != nil {
		return RespGetUnionOrderLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUnionOrderLicense{}, bizErr
	}
	return resp, nil
}
