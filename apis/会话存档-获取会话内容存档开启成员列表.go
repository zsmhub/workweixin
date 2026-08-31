package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetPermitUserListMsgaudit 获取会话内容存档开启成员列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/91614#获取会话内容存档开启成员列表
type ReqGetPermitUserListMsgaudit struct {
	// Type 拉取对应版本的开启成员列表。1表示办公版；2表示服务版；3表示企业版。非必填，不填写的时候返回全量成员列表。
	Type int `json:"type,omitempty"`
}

var _ bodyer = ReqGetPermitUserListMsgaudit{}

func (x ReqGetPermitUserListMsgaudit) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetPermitUserListMsgaudit 获取会话内容存档开启成员列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/91614#获取会话内容存档开启成员列表
type RespGetPermitUserListMsgaudit struct {
	CommonResp
	// Ids 设置在开启范围内的成员的userid列表
	Ids []string `json:"ids"`
}

var _ bodyer = RespGetPermitUserListMsgaudit{}

func (x RespGetPermitUserListMsgaudit) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetPermitUserListMsgaudit 获取会话内容存档开启成员列表
// 文档：https://developer.work.weixin.qq.com/document/path/91614#获取会话内容存档开启成员列表
func (c *ApiClient) ExecGetPermitUserListMsgaudit(req ReqGetPermitUserListMsgaudit) (RespGetPermitUserListMsgaudit, error) {
	var resp RespGetPermitUserListMsgaudit
	err := c.executeWXApiPost("/cgi-bin/msgaudit/get_permit_user_list", req, &resp, true)
	if err != nil {
		return RespGetPermitUserListMsgaudit{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetPermitUserListMsgaudit{}, bizErr
	}
	return resp, nil
}
