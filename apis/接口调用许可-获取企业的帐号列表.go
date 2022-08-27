package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListActivedAccountLicense 获取企业的帐号列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/95544#获取企业的帐号列表
type ReqListActivedAccountLicense struct {
	// Corpid 企业corpid ，只支持加密的corpid，必填
	Corpid string `json:"corpid"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500
	Limit int `json:"limit,omitempty"`
}

var _ bodyer = ReqListActivedAccountLicense{}

func (x ReqListActivedAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListActivedAccountLicense 获取企业的帐号列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/95544#获取企业的帐号列表
type RespListActivedAccountLicense struct {
	AccountList []struct {
		// ActiveTime 激活时间
		ActiveTime int `json:"active_time"`
		// ExpireTime 过期时间
		ExpireTime int `json:"expire_time"`
		// Type 激活码帐号类型:1:基础帐号，2:互通帐号
		Type int `json:"type"`
		// Userid 企业的成员userid。返回加密的userid
		Userid string `json:"userid"`
	} `json:"account_list"` // 已激活成员列表，已激活过期的也会返回
	CommonResp
	// HasMore 是否结束
	HasMore int `json:"has_more"`
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListActivedAccountLicense{}

func (x RespListActivedAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListActivedAccountLicense 获取企业的帐号列表
// 文档：https://developer.work.weixin.qq.com/document/path/95544#获取企业的帐号列表
func (c *ApiClient) ExecListActivedAccountLicense(req ReqListActivedAccountLicense) (RespListActivedAccountLicense, error) {
	var resp RespListActivedAccountLicense
	err := c.executeWXApiPost("/cgi-bin/license/list_actived_account", req, &resp, true)
	if err != nil {
		return RespListActivedAccountLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListActivedAccountLicense{}, bizErr
	}
	return resp, nil
}
