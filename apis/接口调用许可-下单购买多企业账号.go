package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateNewOrderJobLicense 创建多企业新购任务请求
// 文档：https://developer.work.weixin.qq.com/document/path/98892#创建多企业新购任务
type ReqCreateNewOrderJobLicense struct {
	BuyList []struct {
		AccountCount struct {
			// BaseCount 基础帐号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
			BaseCount int `json:"base_count,omitempty"`
			// ExternalContactCount 互通帐号个数，最多1000000个。(若企业为服务商测试企业，最多购买1000个)
			ExternalContactCount int `json:"external_contact_count,omitempty"`
		} `json:"account_count"` // 账号个数详情，基础账号跟互通账号不能同时为0，必填
		AccountDuration struct {
			// Days 购买的天数
			Days int `json:"days,omitempty"`
			// Months 购买的月数，每个月按照31天计算
			Months int `json:"months,omitempty"`
		} `json:"account_duration"` // 帐号购买时长。总购买时长为(months*31+days)天，最少购买1个月(31天)，最多购买60个月(1860天)，必填
		// AutoActiveStatus 是否开启自动激活，不填<strong>默认开启</strong>。<br/>0:关闭 ，1:开启
		AutoActiveStatus int `json:"auto_active_status,omitempty"`
		// Corpid 企业id，必填
		Corpid string `json:"corpid"`
	} `json:"buy_list"` // 企业新购信息列表，详见<strong>BuyInfo</strong>。每次最多传10个，每个jobid最多关联100000个<strong>BuyInfo</strong>，必填
	// Jobid 多企业新购任务id:<br/><strong>不传</strong>默认创建一个新任务。<br/><strong>有传</strong>必须为第一次调用后返回的jobid，可以通过该接口将该任务关联多个新企业的购买账号信息
	Jobid string `json:"jobid,omitempty"`
}

var _ bodyer = ReqCreateNewOrderJobLicense{}

func (x ReqCreateNewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateNewOrderJobLicense 创建多企业新购任务响应
// 文档：https://developer.work.weixin.qq.com/document/path/98892#创建多企业新购任务
type RespCreateNewOrderJobLicense struct {
	CommonResp
	InvalidList []struct {
		// Corpid 不合法的企业 ID
		Corpid string `json:"corpid"`
		CommonResp
	} `json:"invalid_list"` // 不合法的新购信息列表，详见<strong>InvalidInfo</strong>
	// Jobid 多企业新购任务id，当请求包中未指定jobid时，会生成一个新的jobid返回
	Jobid string `json:"jobid"`
}

var _ bodyer = RespCreateNewOrderJobLicense{}

func (x RespCreateNewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateNewOrderJobLicense 创建多企业新购任务
// 文档：https://developer.work.weixin.qq.com/document/path/98892#创建多企业新购任务
func (c *ApiClient) ExecCreateNewOrderJobLicense(req ReqCreateNewOrderJobLicense) (RespCreateNewOrderJobLicense, error) {
	var resp RespCreateNewOrderJobLicense
	err := c.executeWXApiPost("/cgi-bin/license/create_new_order_job", req, &resp, true)
	if err != nil {
		return RespCreateNewOrderJobLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateNewOrderJobLicense{}, bizErr
	}
	return resp, nil
}

// ReqSubmitNewOrderJobLicense 提交多企业新购订单请求
// 文档：https://developer.work.weixin.qq.com/document/path/98892#提交多企业新购订单
type ReqSubmitNewOrderJobLicense struct {
	// BuyerUserid 下单人。服务商企业内成员userid。该userid必须登录过企业微信，并且企业微信已绑定微信，必填
	BuyerUserid string `json:"buyer_userid"`
	// Jobid 多企业新购任务id，必填
	Jobid string `json:"jobid"`
}

var _ bodyer = ReqSubmitNewOrderJobLicense{}

func (x ReqSubmitNewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespSubmitNewOrderJobLicense 提交多企业新购订单响应
// 文档：https://developer.work.weixin.qq.com/document/path/98892#提交多企业新购订单
type RespSubmitNewOrderJobLicense struct {
	CommonResp
	// Jobid 多企业新购任务id
	Jobid string `json:"jobid"`
}

var _ bodyer = RespSubmitNewOrderJobLicense{}

func (x RespSubmitNewOrderJobLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSubmitNewOrderJobLicense 提交多企业新购订单
// 文档：https://developer.work.weixin.qq.com/document/path/98892#提交多企业新购订单
func (c *ApiClient) ExecSubmitNewOrderJobLicense(req ReqSubmitNewOrderJobLicense) (RespSubmitNewOrderJobLicense, error) {
	var resp RespSubmitNewOrderJobLicense
	err := c.executeWXApiPost("/cgi-bin/license/submit_new_order_job", req, &resp, true)
	if err != nil {
		return RespSubmitNewOrderJobLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSubmitNewOrderJobLicense{}, bizErr
	}
	return resp, nil
}

// ReqNewOrderJobResultLicense 获取多企业新购订单提交结果请求
// 文档：https://developer.work.weixin.qq.com/document/path/98892#获取多企业新购订单提交结果
type ReqNewOrderJobResultLicense struct {
	// Jobid 多企业新购任务id，必填
	Jobid string `json:"jobid"`
}

var _ bodyer = ReqNewOrderJobResultLicense{}

func (x ReqNewOrderJobResultLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespNewOrderJobResultLicense 获取多企业新购订单提交结果响应
// 文档：https://developer.work.weixin.qq.com/document/path/98892#获取多企业新购订单提交结果
type RespNewOrderJobResultLicense struct {
	CommonResp
	FailList []struct {
		// Corpid 下单失败的企业 ID
		Corpid string `json:"corpid"`
		CommonResp
	} `json:"fail_list"` // 下单失败的企业及原因。详见<strong>FailInfo</strong>
	// OrderID 订单号。创建完成后返回
	OrderID string `json:"order_id"`
	// Status 订单创建结果。<br/> 1:创建完成 <br/>2:创建中，稍后再试<br/>3:创建失败
	Status int `json:"status"`
}

var _ bodyer = RespNewOrderJobResultLicense{}

func (x RespNewOrderJobResultLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecNewOrderJobResultLicense 获取多企业新购订单提交结果
// 文档：https://developer.work.weixin.qq.com/document/path/98892#获取多企业新购订单提交结果
func (c *ApiClient) ExecNewOrderJobResultLicense(req ReqNewOrderJobResultLicense) (RespNewOrderJobResultLicense, error) {
	var resp RespNewOrderJobResultLicense
	err := c.executeWXApiPost("/cgi-bin/license/new_order_job_result", req, &resp, true)
	if err != nil {
		return RespNewOrderJobResultLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespNewOrderJobResultLicense{}, bizErr
	}
	return resp, nil
}
