package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListOrderAccountLicense 获取订单中的帐号列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/95649#获取订单中的帐号列表
type ReqListOrderAccountLicense struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500
	Limit int `json:"limit,omitempty"`
	// OrderID 订单号，必填
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqListOrderAccountLicense{}

func (x ReqListOrderAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListOrderAccountLicense 获取订单中的帐号列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/95649#获取订单中的帐号列表
type RespListOrderAccountLicense struct {
	AccountList []struct {
		// ActiveCode 帐号码，订单类型为购买帐号时，返回该字段
		ActiveCode string `json:"active_code"`
		// Type 帐号类型:1:基础帐号，2:互通帐号
		Type int `json:"type"`
		// Userid 企业续期成员userid，订单类型为续期帐号时，返回该字段。返回加密的userid
		Userid string `json:"userid"`
	} `json:"account_list"` // 帐号列表
	CommonResp
	// HasMore 是否结束
	HasMore int `json:"has_more"`
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListOrderAccountLicense{}

func (x RespListOrderAccountLicense) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListOrderAccountLicense 获取订单中的帐号列表
// 文档：https://developer.work.weixin.qq.com/document/path/95649#获取订单中的帐号列表
func (c *ApiClient) ExecListOrderAccountLicense(req ReqListOrderAccountLicense) (RespListOrderAccountLicense, error) {
	var resp RespListOrderAccountLicense
	err := c.executeWXApiPost("/cgi-bin/license/list_order_account", req, &resp, true)
	if err != nil {
		return RespListOrderAccountLicense{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListOrderAccountLicense{}, bizErr
	}
	return resp, nil
}
