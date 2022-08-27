package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqTransferCustomerResigned 分配离职成员的客户请求
// 文档：https://developer.work.weixin.qq.com/document/path/94100#分配离职成员的客户
type ReqTransferCustomerResigned struct {
	// ExternalUserid 客户的external_userid列表，最多一次转移100个客户，必填
	ExternalUserid []string `json:"external_userid"`
	// HandoverUserid 原跟进成员的userid，必填
	HandoverUserid string `json:"handover_userid"`
	// TakeoverUserid 接替成员的userid，必填
	TakeoverUserid string `json:"takeover_userid"`
}

var _ bodyer = ReqTransferCustomerResigned{}

func (x ReqTransferCustomerResigned) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespTransferCustomerResigned 分配离职成员的客户响应
// 文档：https://developer.work.weixin.qq.com/document/path/94100#分配离职成员的客户
type RespTransferCustomerResigned struct {
	Customer []struct {
		// Errcode 对此客户进行分配的结果, 具体可参考<a href="#10649" rel="nofollow">全局错误码</a>, <strong>0表示开始分配流程,待24小时后自动接替,并不代表最终分配成功</strong>
		Errcode int `json:"errcode"`
		// ExternalUserid 客户的<code>external_userid</code>
		ExternalUserid string `json:"external_userid"`
	} `json:"customer"`
	// Errcode 对此客户进行分配的结果, 具体可参考<a href="#10649" rel="nofollow">全局错误码</a>, <strong>0表示开始分配流程,待24小时后自动接替,并不代表最终分配成功</strong>
	CommonResp
}

var _ bodyer = RespTransferCustomerResigned{}

func (x RespTransferCustomerResigned) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecTransferCustomerResigned 离职继承-分配离职成员的客户
// 文档：https://developer.work.weixin.qq.com/document/path/94100#分配离职成员的客户
func (c *ApiClient) ExecTransferCustomerResigned(req ReqTransferCustomerResigned) (RespTransferCustomerResigned, error) {
	var resp RespTransferCustomerResigned
	err := c.executeWXApiPost("/cgi-bin/externalcontact/resigned/transfer_customer", req, &resp, true)
	if err != nil {
		return RespTransferCustomerResigned{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespTransferCustomerResigned{}, bizErr
	}
	return resp, nil
}
