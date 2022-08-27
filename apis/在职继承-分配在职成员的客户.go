package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqTransferCustomerExternalcontact 分配在职成员的客户请求
// 文档：https://developer.work.weixin.qq.com/document/path/94096#分配在职成员的客户
type ReqTransferCustomerExternalcontact struct {
	// ExternalUserid 客户的external_userid列表，每次最多分配100个客户，必填
	ExternalUserid []string `json:"external_userid"`
	// HandoverUserid 原跟进成员的userid，必填
	HandoverUserid string `json:"handover_userid"`
	// TakeoverUserid 接替成员的userid，必填
	TakeoverUserid string `json:"takeover_userid"`
	// TransferSuccessMsg 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案
	TransferSuccessMsg string `json:"transfer_success_msg"`
}

var _ bodyer = ReqTransferCustomerExternalcontact{}

func (x ReqTransferCustomerExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespTransferCustomerExternalcontact 分配在职成员的客户响应
// 文档：https://developer.work.weixin.qq.com/document/path/94096#分配在职成员的客户
type RespTransferCustomerExternalcontact struct {
	Customer []struct {
		// Errcode 对此客户进行分配的结果, 具体可参考<a href="#10649" rel="nofollow">全局错误码</a>, <strong>0表示成功发起接替,待24小时后自动接替,并不代表最终接替成功</strong>
		Errcode int `json:"errcode"`
		// ExternalUserid 客户的<code>external_userid</code>
		ExternalUserid string `json:"external_userid"`
	} `json:"customer"`
	// Errcode 对此客户进行分配的结果, 具体可参考<a href="#10649" rel="nofollow">全局错误码</a>, <strong>0表示成功发起接替,待24小时后自动接替,并不代表最终接替成功</strong>
	CommonResp
}

var _ bodyer = RespTransferCustomerExternalcontact{}

func (x RespTransferCustomerExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecTransferCustomerExternalcontact 在职继承-分配在职成员的客户
// 文档：https://developer.work.weixin.qq.com/document/path/94096#分配在职成员的客户
func (c *ApiClient) ExecTransferCustomerExternalcontact(req ReqTransferCustomerExternalcontact) (RespTransferCustomerExternalcontact, error) {
	var resp RespTransferCustomerExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/transfer_customer", req, &resp, true)
	if err != nil {
		return RespTransferCustomerExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespTransferCustomerExternalcontact{}, bizErr
	}
	return resp, nil
}
