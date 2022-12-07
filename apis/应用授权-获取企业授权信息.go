package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetAuthInfoService 获取企业授权信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/90604#获取企业授权信息
type ReqGetAuthInfoService struct {
	// AuthCorpid 授权方corpid，必填
	AuthCorpid string `json:"auth_corpid"`
	// PermanentCode 永久授权码，通过get_permanent_code获取，必填
	PermanentCode string `json:"permanent_code"`
}

var _ bodyer = ReqGetAuthInfoService{}

func (x ReqGetAuthInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetAuthInfoService 获取企业授权信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/90604#获取企业授权信息
type RespGetAuthInfoService struct {
	CommonResp
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       AuthInfo       `json:"auth_info"`
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	// 收银台获取企业授权信息额外返回字段：https://developer.work.weixin.qq.com/document/path/91912
	EditionInfo EditionInfo `json:"edition_info"`
}

var _ bodyer = RespGetAuthInfoService{}

func (x RespGetAuthInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetAuthInfoService 获取企业授权信息
// 文档：https://developer.work.weixin.qq.com/document/path/90604#获取企业授权信息
func (c *ApiClient) ExecGetAuthInfoService(req ReqGetAuthInfoService) (RespGetAuthInfoService, error) {
	var resp RespGetAuthInfoService
	err := c.executeWXApiPost("/cgi-bin/service/get_auth_info", req, &resp, true)
	if err != nil {
		return RespGetAuthInfoService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAuthInfoService{}, bizErr
	}
	return resp, nil
}
