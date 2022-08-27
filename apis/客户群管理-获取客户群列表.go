package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqListGroupchat 获取客户群列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/93414#获取客户群列表
type (
	GroupchatOwnerFilter struct {
		// UseridList 用户ID列表。最多100个
		UseridList []string `json:"userid_list"`
	}
)

type ReqListGroupchat struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 1000，必填
	Limit       int                  `json:"limit"`
	OwnerFilter GroupchatOwnerFilter `json:"owner_filter"` // 群主过滤。<br/>如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
	// StatusFilter 客户群跟进状态过滤。<br/>0 - 所有列表(即不过滤)<br/>1 - 离职待继承<br/>2 - 离职继承中<br/>3 - 离职继承完成<br/><br/>默认为0
	StatusFilter int `json:"status_filter"`
}

var _ bodyer = ReqListGroupchat{}

func (x ReqListGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListGroupchat 获取客户群列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/93414#获取客户群列表
type RespListGroupchat struct {
	CommonResp
	GroupChatList []struct {
		ChatID string `json:"chat_id"`
		Status int    `json:"status"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListGroupchat{}

func (x RespListGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListGroupchat 获取客户群列表
// 文档：https://developer.work.weixin.qq.com/document/path/93414#获取客户群列表
func (c *ApiClient) ExecListGroupchat(req ReqListGroupchat) (RespListGroupchat, error) {
	var resp RespListGroupchat
	err := c.executeWXApiPost("/cgi-bin/externalcontact/groupchat/list", req, &resp, true)
	if err != nil {
		return RespListGroupchat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListGroupchat{}, bizErr
	}
	return resp, nil
}
