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
	AuthCorpInfo struct {
		CorpFullName      string `json:"corp_full_name"`
		CorpIndustry      string `json:"corp_industry"`
		CorpName          string `json:"corp_name"`
		CorpScale         string `json:"corp_scale"`
		CorpSquareLogoURL string `json:"corp_square_logo_url"`
		CorpSubIndustry   string `json:"corp_sub_industry"`
		CorpType          string `json:"corp_type"`
		CorpUserMax       int    `json:"corp_user_max"`
		CorpWxqrcode      string `json:"corp_wxqrcode"`
		Corpid            string `json:"corpid"`
		SubjectType       int    `json:"subject_type"`
		VerifiedEndTime   int    `json:"verified_end_time"`
		Location          string `json:"location"`
	} `json:"auth_corp_info"`
	AuthInfo struct {
		Agent []struct {
			Agentid         int    `json:"agentid"`
			Appid           int    `json:"appid"`
			AuthMode        int    `json:"auth_mode"`
			IsCustomizedApp bool   `json:"is_customized_app"`
			Name            string `json:"name"`
			Privilege       struct {
				AllowParty []int    `json:"allow_party"`
				AllowTag   []int    `json:"allow_tag"`
				AllowUser  []string `json:"allow_user"`
				ExtraParty []int    `json:"extra_party"`
				ExtraTag   []int    `json:"extra_tag"`
				ExtraUser  []string `json:"extra_user"`
				Level      int      `json:"level"`
			} `json:"privilege"`
			RoundLogoURL string `json:"round_logo_url"`
			SharedFrom   struct {
				Corpid string `json:"corpid"`
			} `json:"shared_from"`
			SquareLogoURL string `json:"square_logo_url"`
		} `json:"agent"`
	} `json:"auth_info"`
	DealerCorpInfo struct {
		CorpName string `json:"corp_name"`
		Corpid   string `json:"corpid"`
	} `json:"dealer_corp_info"`
	CommonResp
}

var _ bodyer = RespGetAuthInfoService{}

func (x RespGetAuthInfoService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// execGetAuthInfoService 获取企业授权信息
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
