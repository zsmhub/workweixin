package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqGetSuiteTokenService 获取第三方应用凭证请求
// 文档：https://developer.work.weixin.qq.com/document/path/90600#获取第三方应用凭证
type ReqGetSuiteTokenService struct {
	// SuiteID 以ww或wx开头应用id（对应于旧的以tj开头的套件id），必填
	SuiteID string `json:"suite_id"`
	// SuiteSecret 应用secret，必填
	SuiteSecret string `json:"suite_secret"`
	// SuiteTicket 企业微信后台推送的ticket，必填
	SuiteTicket string `json:"suite_ticket"`
}

var _ bodyer = ReqGetSuiteTokenService{}

func (x ReqGetSuiteTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetSuiteTokenService 获取第三方应用凭证响应
// 文档：https://developer.work.weixin.qq.com/document/path/90600#获取第三方应用凭证
type RespGetSuiteTokenService struct {
	CommonResp
	ExpiresIn        int    `json:"expires_in"`
	SuiteAccessToken string `json:"suite_access_token"`
}

var _ bodyer = RespGetSuiteTokenService{}

func (x RespGetSuiteTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetSuiteTokenService 获取第三方应用凭证
// 文档：https://developer.work.weixin.qq.com/document/path/90600#获取第三方应用凭证
func (c *ApiClient) ExecGetSuiteTokenService(req ReqGetSuiteTokenService) (RespGetSuiteTokenService, error) {
	var resp RespGetSuiteTokenService
	err := c.executeWXApiPost("/cgi-bin/service/get_suite_token", req, &resp, false)
	if err != nil {
		return RespGetSuiteTokenService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetSuiteTokenService{}, bizErr
	}
	return resp, nil
}
