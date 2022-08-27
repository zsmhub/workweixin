package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqInviteBatch 成邀请成员请求
// 文档：https://developer.work.weixin.qq.com/document/path/91127#成邀请成员
type ReqInviteBatch struct {
	User  []string `json:"user"`
	Party []int    `json:"party"`
	Tag   []int    `json:"tag"`
}

var _ bodyer = ReqInviteBatch{}

func (x ReqInviteBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespInviteBatch 成邀请成员响应
// 文档：https://developer.work.weixin.qq.com/document/path/91127#成邀请成员
type RespInviteBatch struct {
	CommonResp
	// 非法成员列表
	InvalidUser []string `json:"invaliduser"`
	// 非法部门列表
	InvalidParty []int `json:"invalidparty"`
	// 非法标签列表
	InvalidTag []int `json:"invalidtag"`
}

var _ bodyer = RespInviteBatch{}

func (x RespInviteBatch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecInviteBatch 成邀请成员
// 文档：https://developer.work.weixin.qq.com/document/path/91127#成邀请成员
func (c *ApiClient) ExecInviteBatch(req ReqInviteBatch) (RespInviteBatch, error) {
	var resp RespInviteBatch
	err := c.executeWXApiPost("/cgi-bin/batch/invite", req, &resp, true)
	if err != nil {
		return RespInviteBatch{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespInviteBatch{}, bizErr
	}
	return resp, nil
}
