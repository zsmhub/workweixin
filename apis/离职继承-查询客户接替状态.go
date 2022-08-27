package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqTransferResultResigned 查询客户接替状态请求
// 文档：https://developer.work.weixin.qq.com/document/path/94101#查询客户接替状态
type ReqTransferResultResigned struct {
	// Cursor 分页查询的cursor，每个分页返回的数据不会超过1000条；不填或为空表示获取第一个分页
	Cursor string `json:"cursor"`
	// HandoverUserid 原添加成员的<code>userid</code>，必填
	HandoverUserid string `json:"handover_userid"`
	// TakeoverUserid 接替成员的<code>userid</code>，必填
	TakeoverUserid string `json:"takeover_userid"`
}

var _ bodyer = ReqTransferResultResigned{}

func (x ReqTransferResultResigned) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespTransferResultResigned 查询客户接替状态响应
// 文档：https://developer.work.weixin.qq.com/document/path/94101#查询客户接替状态
type RespTransferResultResigned struct {
	Customer []struct {
		// ExternalUserid 转接客户的<code>外部联系人userid</code>
		ExternalUserid string `json:"external_userid"`
		// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限
		Status int `json:"status"`
		// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
		TakeoverTime int `json:"takeover_time"`
	} `json:"customer"`
	CommonResp
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespTransferResultResigned{}

func (x RespTransferResultResigned) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecTransferResultResigned 离职继承-查询客户接替状态
// 文档：https://developer.work.weixin.qq.com/document/path/94101#查询客户接替状态
func (c *ApiClient) ExecTransferResultResigned(req ReqTransferResultResigned) (RespTransferResultResigned, error) {
	var resp RespTransferResultResigned
	err := c.executeWXApiPost("/cgi-bin/externalcontact/resigned/transfer_result", req, &resp, true)
	if err != nil {
		return RespTransferResultResigned{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespTransferResultResigned{}, bizErr
	}
	return resp, nil
}
