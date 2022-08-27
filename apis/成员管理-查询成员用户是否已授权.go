package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCheckMemberAuthUser 查询成员用户是否已授权请求
// 文档：https://developer.work.weixin.qq.com/document/path/94514#查询成员用户是否已授权
type ReqCheckMemberAuthUser struct {
	// OpenUserid  
	OpenUserid string `json:"open_userid,omitempty"`
}

var _ bodyer = ReqCheckMemberAuthUser{}

func (x ReqCheckMemberAuthUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCheckMemberAuthUser 查询成员用户是否已授权响应
// 文档：https://developer.work.weixin.qq.com/document/path/94514#查询成员用户是否已授权
type RespCheckMemberAuthUser struct {
	CommonResp
	// IsMemberAuth <a href="#30245" rel="nofollow">成员授权模式</a>下，该成员已授权
	IsMemberAuth bool `json:"is_member_auth"`
}

var _ bodyer = RespCheckMemberAuthUser{}

func (x RespCheckMemberAuthUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCheckMemberAuthUser 查询成员用户是否已授权
// 文档：https://developer.work.weixin.qq.com/document/path/94514#查询成员用户是否已授权
func (c *ApiClient) ExecCheckMemberAuthUser(req ReqCheckMemberAuthUser) (RespCheckMemberAuthUser, error) {
	var resp RespCheckMemberAuthUser
	err := c.executeWXApiPost("/cgi-bin/user/check_member_auth", req, &resp, true)
	if err != nil {
		return RespCheckMemberAuthUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCheckMemberAuthUser{}, bizErr
	}
	return resp, nil
}
