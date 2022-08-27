package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqTransferGroupchat 分配离职成员的客户群请求
// 文档：https://developer.work.weixin.qq.com/document/path/93242#分配离职成员的客户群
type ReqTransferGroupchat struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围: 1 ~ 100，必填
	ChatIDList []string `json:"chat_id_list"`
	// NewOwner 新群主ID，必填
	NewOwner string `json:"new_owner"`
}

var _ bodyer = ReqTransferGroupchat{}

func (x ReqTransferGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespTransferGroupchat 分配离职成员的客户群响应
// 文档：https://developer.work.weixin.qq.com/document/path/93242#分配离职成员的客户群
type RespTransferGroupchat struct {
	// Errcode 没能成功继承的群，错误码
	CommonResp
	Errmsg         string `json:"errmsg"`
	FailedChatList []struct {
		// ChatID 没能成功继承的群ID
		ChatID string `json:"chat_id"`
		// Errcode 没能成功继承的群，错误码
		CommonResp
		Errmsg string `json:"errmsg"`
	} `json:"failed_chat_list"` // 没能成功继承的群
}

var _ bodyer = RespTransferGroupchat{}

func (x RespTransferGroupchat) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecTransferGroupchat 离职继承-分配离职成员的客户群
// 文档：https://developer.work.weixin.qq.com/document/path/93242#分配离职成员的客户群
func (c *ApiClient) ExecTransferGroupchat(req ReqTransferGroupchat) (RespTransferGroupchat, error) {
	var resp RespTransferGroupchat
	err := c.executeWXApiPost("/cgi-bin/externalcontact/groupchat/transfer", req, &resp, true)
	if err != nil {
		return RespTransferGroupchat{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespTransferGroupchat{}, bizErr
	}
	return resp, nil
}
