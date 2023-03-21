package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetOrderListPaytool 获取收款订单列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/98053#获取收款订单列表
type ReqGetOrderListPaytool struct {
	// BusinessType 业务类型
	BusinessType int `json:"business_type,omitempty"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor,omitempty"`
	// EndTime 结束时间
	EndTime int `json:"end_time,omitempty"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 2000
	Limit int `json:"limit,omitempty"`
	// NonceStr 随机字符串，长度要求在32字节以内，用于保证签名不可预测及防重放攻击。<br/>需保证15分钟内不能重复，推荐随机字符串生成算法
	NonceStr string `json:"nonce_str,omitempty"`
	// Sig 数字签名
	Sig string `json:"sig,omitempty"`
	// StartTime 起始时间
	StartTime int `json:"start_time,omitempty"`
	// Ts unix时间戳（中国时区）,精确到秒。<br/>注意业务系统的机器时间与腾讯的时间相差不能超过15分钟
	Ts int `json:"ts,omitempty"`
}

var _ bodyer = ReqGetOrderListPaytool{}

func (x ReqGetOrderListPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetOrderListPaytool 获取收款订单列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/98053#获取收款订单列表
type RespGetOrderListPaytool struct {
	CommonResp
	// HasMore 是否还有更多数据
	HasMore int `json:"has_more"`
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录
	NextCursor   string `json:"next_cursor"`
	PayOrderList []struct {
		// BuyContent 购买内容
		BuyContent string `json:"buy_content"`
		// CreateTime 订单创建时间
		CreateTime int `json:"create_time"`
		// Creator 订单创建人
		Creator string `json:"creator"`
		// CustomCorpid 客户企业的corpid
		CustomCorpid string `json:"custom_corpid"`
		// OrderFrom 订单来源。取值范围为:<br/> 1 - 客户下单<br/> 2 - 服务商创建
		OrderFrom int `json:"order_from"`
		// OrderID 订单号
		OrderID string `json:"order_id"`
		// OrderStatus 订单状态。取值范围为:<br/>1 - 待支付 <br/> 2 - 已支付<br/> 3 - 订单取消<br/> 4 - 支付过期<br/> 5 - 退款申请中<br/> 6 - 已退款<br/> 7 - 交易完成
		OrderStatus int `json:"order_status"`
		// OriginPrice 原价金额
		OriginPrice int `json:"origin_price"`
		// PaidPrice 实付金额
		PaidPrice int `json:"paid_price"`
	} `json:"pay_order_list"` // 订单列表
}

var _ bodyer = RespGetOrderListPaytool{}

func (x RespGetOrderListPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetOrderListPaytool 获取收款订单列表
// 文档：https://developer.work.weixin.qq.com/document/path/98053#获取收款订单列表
func (c *ApiClient) ExecGetOrderListPaytool(req ReqGetOrderListPaytool) (RespGetOrderListPaytool, error) {
	var resp RespGetOrderListPaytool
	err := c.executeWXApiPost("/cgi-bin/paytool/get_order_list", req, &resp, true)
	if err != nil {
		return RespGetOrderListPaytool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetOrderListPaytool{}, bizErr
	}
	return resp, nil
}
