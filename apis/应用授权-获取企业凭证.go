package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

const PathGetCorpToken = "/cgi-bin/service/get_corp_token"

// ReqGetCorpTokenService 获取企业凭证请求
// 文档：https://developer.work.weixin.qq.com/document/path/90605#获取企业凭证
type ReqGetCorpTokenService struct {
	// AuthCorpid 授权方corpid，必填
	AuthCorpid string `json:"auth_corpid"`
	// PermanentCode 永久授权码，通过get_permanent_code获取，必填
	PermanentCode string `json:"permanent_code"`
}

var _ bodyer = ReqGetCorpTokenService{}

func (x ReqGetCorpTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetCorpTokenService 获取企业凭证响应
// 文档：https://developer.work.weixin.qq.com/document/path/90605#获取企业凭证
type RespGetCorpTokenService struct {
	AccessToken string `json:"access_token"`
	CommonResp
	ExpiresIn int `json:"expires_in"`
}

var _ bodyer = RespGetCorpTokenService{}

func (x RespGetCorpTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetCorpTokenService 获取企业凭证
// 文档：https://developer.work.weixin.qq.com/document/path/90605#获取企业凭证
func (c *ApiClient) ExecGetCorpTokenService(req ReqGetCorpTokenService) (RespGetCorpTokenService, error) {
	var resp RespGetCorpTokenService
	// 注意：此接口是授权企业客户端调用的，但是需要的 access_token 是第三方应用的 suite_access_token
	err := c.executeWXApiPost(PathGetCorpToken, req, &resp, true)
	if err != nil {
		return RespGetCorpTokenService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCorpTokenService{}, bizErr
	}
	return resp, nil
}
