package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqListExternalcontact 获取客户列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/92264#获取客户列表
type ReqListExternalcontact struct {
	// Userid 企业成员的userid，必填
	Userid string `json:"userid"`
}

var _ urlValuer = ReqListExternalcontact{}

func (x ReqListExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListExternalcontact 获取客户列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/92264#获取客户列表
type RespListExternalcontact struct {
	CommonResp
	ExternalUserid []string `json:"external_userid"`
}

var _ bodyer = RespListExternalcontact{}

func (x RespListExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListExternalcontact 获取客户列表
// 文档：https://developer.work.weixin.qq.com/document/path/92264#获取客户列表
func (c *ApiClient) ExecListExternalcontact(req ReqListExternalcontact) (RespListExternalcontact, error) {
	var resp RespListExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/list", req, &resp, true)
	if err != nil {
		return RespListExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListExternalcontact{}, bizErr
	}
	return resp, nil
}
