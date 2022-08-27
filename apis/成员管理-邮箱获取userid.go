package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetUseridByEmailUser 邮箱获取userid请求
// 文档：https://developer.work.weixin.qq.com/document/path/95892#邮箱获取userid
type ReqGetUseridByEmailUser struct {
	// Email 邮箱，必填
	Email string `json:"email"`
	// EmailType 邮箱类型:1-企业邮箱（默认）；2-个人邮箱
	EmailType int `json:"email_type,omitempty"`
}

var _ bodyer = ReqGetUseridByEmailUser{}

func (x ReqGetUseridByEmailUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetUseridByEmailUser 邮箱获取userid响应
// 文档：https://developer.work.weixin.qq.com/document/path/95892#邮箱获取userid
type RespGetUseridByEmailUser struct {
	CommonResp
	// Userid 成员UserID。注意:已升级openid的代开发或第三方，获取的是密文userid
	Userid string `json:"userid"`
}

var _ bodyer = RespGetUseridByEmailUser{}

func (x RespGetUseridByEmailUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetUseridByEmailUser 邮箱获取userid
// 文档：https://developer.work.weixin.qq.com/document/path/95892#邮箱获取userid
func (c *ApiClient) ExecGetUseridByEmailUser(req ReqGetUseridByEmailUser) (RespGetUseridByEmailUser, error) {
	var resp RespGetUseridByEmailUser
	err := c.executeWXApiPost("/cgi-bin/user/get_userid_by_email", req, &resp, true)
	if err != nil {
		return RespGetUseridByEmailUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetUseridByEmailUser{}, bizErr
	}
	return resp, nil
}
