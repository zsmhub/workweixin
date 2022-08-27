package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqAuthsuccUser 成二次验证请求
// 文档：https://developer.work.weixin.qq.com/document/path/90339#二次验证
type ReqAuthsuccUser struct {
	// Userid 成员UserID。对应管理端的帐号，必填
	Userid string `json:"userid"`
}

var _ urlValuer = ReqAuthsuccUser{}

func (x ReqAuthsuccUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespAuthsuccUser 成二次验证响应
// 文档：https://developer.work.weixin.qq.com/document/path/90339#成二次验证
type RespAuthsuccUser struct {
	CommonResp
}

var _ bodyer = RespAuthsuccUser{}

func (x RespAuthsuccUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAuthsuccUser 成二次验证
// 文档：https://developer.work.weixin.qq.com/document/path/90339#成二次验证
func (c *ApiClient) ExecAuthsuccUser(req ReqAuthsuccUser) (RespAuthsuccUser, error) {
	var resp RespAuthsuccUser
	err := c.executeWXApiGet("/cgi-bin/user/authsucc", req, &resp, true)
	if err != nil {
		return RespAuthsuccUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAuthsuccUser{}, bizErr
	}
	return resp, nil
}
