package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetOrderListService 获取订单列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/91910#获取订单列表
type ReqGetOrderListService struct {
	// EndTime 终止时间，UNIX时间戳，必填
	EndTime int `json:"end_time"`
	// StartTime 起始时间，UNIX时间戳，必填
	StartTime int `json:"start_time"`
	// TestMode 指定拉取正式或测试授权的订单。默认值为0，其中 0-正式授权，1-测试授权。
	TestMode int `json:"test_mode,omitempty"`
}

var _ bodyer = ReqGetOrderListService{}

func (x ReqGetOrderListService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetOrderListService 获取订单列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/91910#获取订单列表
type RespGetOrderListService struct {
	CommonResp
	OrderList []struct {
		// Appid 套件应用id（仅旧套件有该字段）
		Appid int `json:"appid"`
		// BeginTime 购买生效期的开始时间（UNIX时间戳）
		BeginTime      int `json:"begin_time"`
		DealerCorpInfo struct {
			// CorpName 代理商的企业简称
			CorpName string `json:"corp_name"`
			// Corpid 代理商corpid
			Corpid string `json:"corpid"`
		} `json:"dealer_corp_info"` // 代理商信息（仅当有渠道商报备后才会有此字段）
		// DealerShareAmount 代理商分成金额，单位分
		DealerShareAmount int `json:"dealer_share_amount"`
		// EditionID 购买版本ID
		EditionID string `json:"edition_id"`
		// EditionName 购买版本名字
		EditionName string `json:"edition_name"`
		// EndTime 购买生效期的结束时间（UNIX时间戳）
		EndTime int `json:"end_time"`
		// OperatorCorpid 下单方corpid
		OperatorCorpid string `json:"operator_corpid"`
		// OperatorID 下单操作人员userid。如果是服务商代下单，没有该字段。
		OperatorID string `json:"operator_id"`
		// OrderFrom 下单来源。0-企业下单；1-服务商代下单；2-代理商代下单
		OrderFrom int `json:"order_from"`
		// OrderPeriod 购买的时长，单位天
		OrderPeriod int `json:"order_period"`
		// OrderStatus 订单状态。订单状态。0-待支付，1-已支付，2-已取消， 3-支付过期， 4-申请退款中， 5-退款成功， 6-退款被拒绝
		OrderStatus int `json:"order_status"`
		// OrderTime 下单时间（UNIX时间戳）
		OrderTime int `json:"order_time"`
		// OrderType 订单类型。0-新购应用，1-扩容应用人数，2-续期应用时间，3-变更版本
		OrderType int `json:"order_type"`
		// Orderid 订单号
		Orderid string `json:"orderid"`
		// PaidCorpid 客户企业的corpid
		PaidCorpid string `json:"paid_corpid"`
		// PaidTime 付款时间（UNIX时间戳）
		PaidTime int `json:"paid_time"`
		// PlatformShareAmount 平台分成金额，单位分
		PlatformShareAmount int `json:"platform_share_amount"`
		// Price 应付价格，单位分
		Price int `json:"price"`
		// ServiceShareAmount 服务商分成金额，单位分
		ServiceShareAmount int `json:"service_share_amount"`
		// Suiteid 应用id
		Suiteid string `json:"suiteid"`
		// UserCount 购买的人数
		UserCount int `json:"user_count"`
	} `json:"order_list"` // 订单列表
}

var _ bodyer = RespGetOrderListService{}

func (x RespGetOrderListService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetOrderListService 获取订单列表
// 文档：https://developer.work.weixin.qq.com/document/path/91910#获取订单列表
func (c *ApiClient) ExecGetOrderListService(req ReqGetOrderListService) (RespGetOrderListService, error) {
	var resp RespGetOrderListService
	err := c.executeWXApiPost("/cgi-bin/service/get_order_list", req, &resp, true)
	if err != nil {
		return RespGetOrderListService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetOrderListService{}, bizErr
	}
	return resp, nil
}
