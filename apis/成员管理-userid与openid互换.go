package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqConvertToOpenidUser userid转openid请求
// 文档：https://developer.work.weixin.qq.com/document/path/90338#userid转openid
type ReqConvertToOpenidUser struct {
	// Userid 企业内的成员id，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqConvertToOpenidUser{}

func (x ReqConvertToOpenidUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespConvertToOpenidUser userid转openid响应
// 文档：https://developer.work.weixin.qq.com/document/path/90338#userid转openid
type RespConvertToOpenidUser struct {
	CommonResp
	// Openid 企业微信成员userid对应的openid
	Openid string `json:"openid"`
}

var _ bodyer = RespConvertToOpenidUser{}

func (x RespConvertToOpenidUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecConvertToOpenidUser userid转openid
// 文档：https://developer.work.weixin.qq.com/document/path/90338#userid转openid
func (c *ApiClient) ExecConvertToOpenidUser(req ReqConvertToOpenidUser) (RespConvertToOpenidUser, error) {
	var resp RespConvertToOpenidUser
	err := c.executeWXApiPost("/cgi-bin/user/convert_to_openid", req, &resp, true)
	if err != nil {
		return RespConvertToOpenidUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespConvertToOpenidUser{}, bizErr
	}
	return resp, nil
}

// ReqConvertToUseridUser openid转userid请求
// 文档：https://developer.work.weixin.qq.com/document/path/90338#openid转userid
type ReqConvertToUseridUser struct {
	// Openid 在使用企业支付之后，返回结果的openid，必填
	Openid string `json:"openid"`
}

var _ bodyer = ReqConvertToUseridUser{}

func (x ReqConvertToUseridUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespConvertToUseridUser openid转userid响应
// 文档：https://developer.work.weixin.qq.com/document/path/90338#openid转userid
type RespConvertToUseridUser struct {
	CommonResp
	// Userid 该openid在企业微信对应的成员userid
	Userid string `json:"userid"`
}

var _ bodyer = RespConvertToUseridUser{}

func (x RespConvertToUseridUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecConvertToUseridUser openid转userid
// 文档：https://developer.work.weixin.qq.com/document/path/90338#openid转userid
func (c *ApiClient) ExecConvertToUseridUser(req ReqConvertToUseridUser) (RespConvertToUseridUser, error) {
	var resp RespConvertToUseridUser
	err := c.executeWXApiPost("/cgi-bin/user/convert_to_userid", req, &resp, true)
	if err != nil {
		return RespConvertToUseridUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespConvertToUseridUser{}, bizErr
	}
	return resp, nil
}
