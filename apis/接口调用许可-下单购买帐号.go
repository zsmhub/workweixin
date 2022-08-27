package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateNewOrderLicense 下单购买帐号请求
// 文档：https://developer.work.weixin.qq.com/document/path/95644#下单购买帐号
type (
	ReqCreateNewOrderLicense struct {
		AccountCount    ReqCreateNewOrderLicenseAccountCount `json:"account_count"`    // 账号个数详情，基础账号跟互通账号不能同时为0，必填
		AccountDuration ReqCreateNewOrderLicenseMonths       `json:"account_duration"` // 帐号购买时长，必填
		// BuyerUserid 下单人。服务商企业内成员userid。该userid必须登录过企业微信，并且企业微信已绑定微信。最终也支持由其他人支付，必填
		BuyerUserid string `json:"buyer_userid"`
		// Corpid 企业id，只支持加密的corpid，必填
		Corpid string `json:"corpid"`
	}

	ReqCreateNewOrderLicenseAccountCount struct {
		// BaseCount 基础帐号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
		BaseCount int `json:"base_count,omitempty"`
		// ExternalContactCount 互通帐号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
		ExternalContactCount int `json:"external_contact_count,omitempty"`
	}

	ReqCreateNewOrderLicenseMonths struct {
		// Months 购买的月数，每个月按照31天计算。最多购买36个月。(若企业为服务商测试企业，最多购买1个月)，必填
		Months int `json:"months"`
	}
)

var _ bodyer = ReqCreateNewOrderLicense{}

func (x ReqCreateNewOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateNewOrderLicense 下单购买帐号响应
// 文档：https://developer.work.weixin.qq.com/document/path/95644#下单购买帐号
type RespCreateNewOrderLicense struct {
	CommonResp
	// OrderID 订单号
	OrderID string `json:"order_id"`
}

var _ bodyer = RespCreateNewOrderLicense{}

func (x RespCreateNewOrderLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateNewOrderLicense 下单购买帐号
// 文档：https://developer.work.weixin.qq.com/document/path/95644#下单购买帐号
func (c *ApiClient) ExecCreateNewOrderLicense(req ReqCreateNewOrderLicense) (RespCreateNewOrderLicense, error) {
	var resp RespCreateNewOrderLicense
	err := c.executeWXApiPost("/cgi-bin/license/create_new_order", req, &resp, true)
	if err != nil {
		return RespCreateNewOrderLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateNewOrderLicense{}, bizErr
	}
	return resp, nil
}
