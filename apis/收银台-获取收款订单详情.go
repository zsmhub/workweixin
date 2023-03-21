package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetOrderDetailPaytool 获取收款订单详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/98054#获取收款订单详情
type ReqGetOrderDetailPaytool struct {
	// NonceStr 随机字符串，长度要求在32字节以内，用于保证签名不可预测及防重放攻击。<br/>需保证15分钟内不能重复，推荐<a href="#47107/%E9%9A%8F%E6%9C%BA%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%94%9F%E6%88%90%E7%AE%97%E6%B3%95" rel="nofollow">随机字符串生成算法</a>
	NonceStr string `json:"nonce_str,omitempty"`
	// OrderID 订单号
	OrderID string `json:"order_id,omitempty"`
	// Sig 数字签名。见<a href="#47107/%E8%AE%A2%E5%8D%95%E7%AD%BE%E5%90%8D%E7%AE%97%E6%B3%95" rel="nofollow">签名算法</a>。签名所需密钥获取路径:<br/>工作台-&gt;企业微信服务商助手-&gt;工具-&gt;收银台-&gt;收银台API调用密钥
	Sig string `json:"sig,omitempty"`
	// Ts unix时间戳（中国时区）,精确到秒。<br/>注意业务系统的机器时间与腾讯的时间相差不能超过15分钟
	Ts int `json:"ts,omitempty"`
}

var _ bodyer = ReqGetOrderDetailPaytool{}

func (x ReqGetOrderDetailPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetOrderDetailPaytool 获取收款订单详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/98054#获取收款订单详情
type RespGetOrderDetailPaytool struct {
	CommonResp
	PayOrder struct {
		OrderId        string `json:"order_id"`
		CreateTime     int    `json:"create_time"`
		CustomCorpid   string `json:"custom_corpid"`
		BuyContent     string `json:"buy_content"`
		OriginPrice    int    `json:"origin_price"`
		PaidPrice      int    `json:"paid_price"`
		OrderStatus    int    `json:"order_status"`
		OrderFrom      int    `json:"order_from"`
		Creator        string `json:"creator"`
		CustomCorpName string `json:"custom_corp_name"`
		PayChannel     int    `json:"pay_channel"`
		ChannelOrderId string `json:"channel_order_id"`
		PaidTime       int    `json:"paid_time"`
		BusinessType   int    `json:"business_type"`
		IncomeType     int    `json:"income_type"`
		IncomeTime     int    `json:"income_time"`
		IncomeAmount   int    `json:"income_amount"`
		ProductList    struct {
			ThirdApp struct {
				OrderType   int `json:"order_type"`
				BuyInfoList []struct {
					Suiteid          string `json:"suiteid"`
					Appid            int    `json:"appid"`
					EditionId        string `json:"edition_id"`
					UserCount        int    `json:"user_count"`
					DurationDays     int    `json:"duration_days"`
					OriginPrice      int    `json:"origin_price"`
					PaidPrice        int    `json:"paid_price"`
					GiftDurationDays int    `json:"gift_duration_days"`
				} `json:"buy_info_list"`
			} `json:"third_app"`
			CustomizedApp struct {
				OrderType   int `json:"order_type"`
				BuyInfoList []struct {
					Suiteid          string `json:"suiteid"`
					UserCount        int    `json:"user_count"`
					DurationDays     int    `json:"duration_days"`
					OriginPrice      int    `json:"origin_price"`
					PaidPrice        int    `json:"paid_price"`
					GiftDurationDays int    `json:"gift_duration_days"`
				} `json:"buy_info_list"`
			} `json:"customized_app"`
			PromotionCase struct {
				OrderType            int    `json:"order_type"`
				CaseId               string `json:"case_id"`
				PromotionEditionName string `json:"promotion_edition_name"`
				BuyInfoList          []struct {
					Suiteid          string `json:"suiteid"`
					Appid            int    `json:"appid"`
					UserCount        int    `json:"user_count"`
					OriginPrice      int    `json:"origin_price"`
					PaidPrice        int    `json:"paid_price"`
					GiftDurationDays int    `json:"gift_duration_days"`
				} `json:"buy_info_list"`
			} `json:"promotion_case"`
		} `json:"product_list"`
	} `json:"pay_order"`
}

var _ bodyer = RespGetOrderDetailPaytool{}

func (x RespGetOrderDetailPaytool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetOrderDetailPaytool 获取收款订单详情
// 文档：https://developer.work.weixin.qq.com/document/path/98054#获取收款订单详情
func (c *ApiClient) ExecGetOrderDetailPaytool(req ReqGetOrderDetailPaytool) (RespGetOrderDetailPaytool, error) {
	var resp RespGetOrderDetailPaytool
	err := c.executeWXApiPost("/cgi-bin/paytool/get_order_detail", req, &resp, true)
	if err != nil {
		return RespGetOrderDetailPaytool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetOrderDetailPaytool{}, bizErr
	}
	return resp, nil
}
