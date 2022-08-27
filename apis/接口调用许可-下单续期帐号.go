package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateRenewOrderJobLicense 创建续期任务请求
// 文档：https://developer.work.weixin.qq.com/document/path/95646#创建续期任务
type (
	ReqCreateRenewOrderJobLicense struct {
		AccountList []ReqCreateRenewOrderJobLicenseAccountListItem `json:"account_list"` // 续期的帐号列表，每次最多1000个。同一个jobid最多关联1000000个基础账号跟1000000个互通账号，必填
		// Corpid 企业id，只支持加密的corpid，必填
		Corpid string `json:"corpid"`
		// Jobid 任务id，若不传则默认创建一个新任务。若指定第一次调用后拿到jobid，可以通过该接口将jobid关联多个userid
		Jobid string `json:"jobid,omitempty"`
	}

	ReqCreateRenewOrderJobLicenseAccountListItem struct {
		// Type 续期帐号类型。1:基础帐号，2:互通帐号，必填
		Type int `json:"type"`
		// Userid 续期企业的成员userid。只支持加密的userid，必填
		Userid string `json:"userid"`
	}
)

var _ bodyer = ReqCreateRenewOrderJobLicense{}

func (x ReqCreateRenewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateRenewOrderJobLicense 创建续期任务响应
// 文档：https://developer.work.weixin.qq.com/document/path/95646#创建续期任务
type RespCreateRenewOrderJobLicense struct {
	// Errcode 账号不合法相关错误码
	CommonResp
	Errmsg             string `json:"errmsg"`
	InvalidAccountList []struct {
		// Errcode 账号不合法相关错误码
		CommonResp
		Errmsg string `json:"errmsg"`
		Type   int    `json:"type"`
		Userid string `json:"userid"`
	} `json:"invalid_account_list"` // 不合法的续期账号列表
	// Jobid 任务id，请求包中未指定jobid时，会生成一个新的jobid返回
	Jobid string `json:"jobid"`
}

var _ bodyer = RespCreateRenewOrderJobLicense{}

func (x RespCreateRenewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateRenewOrderJobLicense 创建续期任务
// 文档：https://developer.work.weixin.qq.com/document/path/95646#创建续期任务
func (c *ApiClient) ExecCreateRenewOrderJobLicense(req ReqCreateRenewOrderJobLicense) (RespCreateRenewOrderJobLicense, error) {
	var resp RespCreateRenewOrderJobLicense
	err := c.executeWXApiPost("/cgi-bin/license/create_renew_order_job", req, &resp, true)
	if err != nil {
		return RespCreateRenewOrderJobLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateRenewOrderJobLicense{}, bizErr
	}
	return resp, nil
}

// ReqSubmitOrderJobLicense 提交续期订单请求
// 文档：https://developer.work.weixin.qq.com/document/path/95646#提交续期订单
type ReqSubmitOrderJobLicense struct {
	AccountDuration struct {
		// Months 购买的月数，每个月按照31天计算。最多购买36个月。(若企业为服务商测试企业，每次续期只能续期1个月)，必填
		Months int `json:"months"`
	} `json:"account_duration"` // 账号购买时长，必填
	// BuyerUserid 下单人。服务商企业内成员userid。该userid必须登录过企业微信，并且企业微信已绑定微信，必填
	BuyerUserid string `json:"buyer_userid"`
	// Jobid 任务id，必填
	Jobid string `json:"jobid"`
}

var _ bodyer = ReqSubmitOrderJobLicense{}

func (x ReqSubmitOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSubmitOrderJobLicense 提交续期订单响应
// 文档：https://developer.work.weixin.qq.com/document/path/95646#提交续期订单
type RespSubmitOrderJobLicense struct {
	CommonResp
	// OrderID 订单号
	OrderID string `json:"order_id"`
}

var _ bodyer = RespSubmitOrderJobLicense{}

func (x RespSubmitOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSubmitOrderJobLicense 提交续期订单
// 文档：https://developer.work.weixin.qq.com/document/path/95646#提交续期订单
func (c *ApiClient) ExecSubmitOrderJobLicense(req ReqSubmitOrderJobLicense) (RespSubmitOrderJobLicense, error) {
	var resp RespSubmitOrderJobLicense
	err := c.executeWXApiPost("/cgi-bin/license/submit_order_job", req, &resp, true)
	if err != nil {
		return RespSubmitOrderJobLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSubmitOrderJobLicense{}, bizErr
	}
	return resp, nil
}
