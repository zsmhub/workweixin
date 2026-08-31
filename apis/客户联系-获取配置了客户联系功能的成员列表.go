package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetFollowUserListExternalcontact 获取配置了客户联系功能的成员列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/92576#获取配置了客户联系功能的成员列表
type ReqGetFollowUserListExternalcontact struct{}

var _ urlValuer = ReqGetFollowUserListExternalcontact{}

func (x ReqGetFollowUserListExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetFollowUserListExternalcontact 获取配置了客户联系功能的成员列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/92576#获取配置了客户联系功能的成员列表
type RespGetFollowUserListExternalcontact struct {
	CommonResp
	// FollowUser 配置了客户联系功能的成员userid列表
	FollowUser []string `json:"follow_user"`
}

var _ bodyer = RespGetFollowUserListExternalcontact{}

func (x RespGetFollowUserListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetFollowUserListExternalcontact 获取配置了客户联系功能的成员列表
// 文档：https://developer.work.weixin.qq.com/document/path/92576#获取配置了客户联系功能的成员列表
func (c *ApiClient) ExecGetFollowUserListExternalcontact(req ReqGetFollowUserListExternalcontact) (RespGetFollowUserListExternalcontact, error) {
	var resp RespGetFollowUserListExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get_follow_user_list", req, &resp, true)
	if err != nil {
		return RespGetFollowUserListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetFollowUserListExternalcontact{}, bizErr
	}
	return resp, nil
}
