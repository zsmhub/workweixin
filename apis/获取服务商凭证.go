package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetProviderTokenService 获取服务商凭证请求
// 文档：https://developer.work.weixin.qq.com/document/path/91200#获取服务商凭证
type ReqGetProviderTokenService struct {
	// Corpid 服务商的corpid，必填
	Corpid string `json:"corpid"`
	// ProviderSecret 服务商的secret，在服务商管理后台可见，必填
	ProviderSecret string `json:"provider_secret"`
}

var _ bodyer = ReqGetProviderTokenService{}

func (x ReqGetProviderTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetProviderTokenService 获取服务商凭证响应
// 文档：https://developer.work.weixin.qq.com/document/path/91200#获取服务商凭证
type RespGetProviderTokenService struct {
	CommonResp

	ExpiresIn           int    `json:"expires_in"`
	ProviderAccessToken string `json:"provider_access_token"`
}

var _ bodyer = RespGetProviderTokenService{}

func (x RespGetProviderTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetProviderTokenService 获取服务商凭证
// 文档：https://developer.work.weixin.qq.com/document/path/91200#获取服务商凭证
func (c *ApiClient) ExecGetProviderTokenService(req ReqGetProviderTokenService) (RespGetProviderTokenService, error) {
	var resp RespGetProviderTokenService
	err := c.executeWXApiPost("/cgi-bin/service/get_provider_token", req, &resp, false)
	if err != nil {
		return RespGetProviderTokenService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetProviderTokenService{}, bizErr
	}
	return resp, nil
}
