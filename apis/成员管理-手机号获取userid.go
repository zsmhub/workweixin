package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuseridUser 手机号获取userid请求
// 文档：https://developer.work.weixin.qq.com/document/path/91693#手机号获取userid
type ReqGetuseridUser struct {
	// Mobile 用户在企业微信通讯录中的手机号码。长度为5~32个字节，必填
	Mobile string `json:"mobile"`
}

var _ bodyer = ReqGetuseridUser{}

func (x ReqGetuseridUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetuseridUser 手机号获取userid响应
// 文档：https://developer.work.weixin.qq.com/document/path/91693#手机号获取userid
type RespGetuseridUser struct {
	CommonResp
	// Userid 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。注意:第三方应用获取的值是密文的userid
	Userid string `json:"userid"`
}

var _ bodyer = RespGetuseridUser{}

func (x RespGetuseridUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetuseridUser 手机号获取userid
// 文档：https://developer.work.weixin.qq.com/document/path/91693#手机号获取userid
func (c *ApiClient) ExecGetuseridUser(req ReqGetuseridUser) (RespGetuseridUser, error) {
	var resp RespGetuseridUser
	err := c.executeWXApiPost("/cgi-bin/user/getuserid", req, &resp, true)
	if err != nil {
		return RespGetuseridUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuseridUser{}, bizErr
	}
	return resp, nil
}
