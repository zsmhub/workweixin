package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqTransferResultExternalcontact 查询客户接替状态请求
// 文档：https://developer.work.weixin.qq.com/document/path/94097#查询客户接替状态
type ReqTransferResultExternalcontact struct {
	// Cursor 分页查询的cursor，每个分页返回的数据不会超过1000条；不填或为空表示获取第一个分页；
	Cursor string `json:"cursor"`
	// HandoverUserid 原添加成员的<code>userid</code>，必填
	HandoverUserid string `json:"handover_userid"`
	// TakeoverUserid 接替成员的<code>userid</code>，必填
	TakeoverUserid string `json:"takeover_userid"`
}

var _ bodyer = ReqTransferResultExternalcontact{}

func (x ReqTransferResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespTransferResultExternalcontact 查询客户接替状态响应
// 文档：https://developer.work.weixin.qq.com/document/path/94097#查询客户接替状态
type RespTransferResultExternalcontact struct {
	Customer []struct {
		// ExternalUserid 转接客户的<code>外部联系人userid</code>
		ExternalUserid string `json:"external_userid"`
		// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
		Status int `json:"status"`
		// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
		TakeoverTime int `json:"takeover_time"`
	} `json:"customer"`
	CommonResp
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespTransferResultExternalcontact{}

func (x RespTransferResultExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecTransferResultExternalcontact 在职继承-查询客户接替状态
// 文档：https://developer.work.weixin.qq.com/document/path/94097#查询客户接替状态
func (c *ApiClient) ExecTransferResultExternalcontact(req ReqTransferResultExternalcontact) (RespTransferResultExternalcontact, error) {
	var resp RespTransferResultExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/transfer_result", req, &resp, true)
	if err != nil {
		return RespTransferResultExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespTransferResultExternalcontact{}, bizErr
	}
	return resp, nil
}
