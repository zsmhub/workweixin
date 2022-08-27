package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetPreAuthCodeService 获取预授权码请求
// 文档：https://developer.work.weixin.qq.com/document/path/90601#获取预授权码
type ReqGetPreAuthCodeService struct{}

var _ urlValuer = ReqGetPreAuthCodeService{}

func (x ReqGetPreAuthCodeService) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetPreAuthCodeService 获取预授权码响应
// 文档：https://developer.work.weixin.qq.com/document/path/90601#获取预授权码
type RespGetPreAuthCodeService struct {
	CommonResp
	// ExpiresIn 有效期（秒）
	ExpiresIn int `json:"expires_in"`
	// PreAuthCode 预授权码,最长为512字节
	PreAuthCode string `json:"pre_auth_code"`
}

var _ bodyer = RespGetPreAuthCodeService{}

func (x RespGetPreAuthCodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetPreAuthCodeService 获取预授权码
// 文档：https://developer.work.weixin.qq.com/document/path/90601#获取预授权码
func (c *ApiClient) ExecGetPreAuthCodeService(req ReqGetPreAuthCodeService) (RespGetPreAuthCodeService, error) {
	var resp RespGetPreAuthCodeService
	err := c.executeWXApiGet("/cgi-bin/service/get_pre_auth_code", req, &resp, true)
	if err != nil {
		return RespGetPreAuthCodeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetPreAuthCodeService{}, bizErr
	}
	return resp, nil
}
