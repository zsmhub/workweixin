package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqGetPermanentCodeService 获取企业永久授权码请求
// 文档：https://developer.work.weixin.qq.com/document/path/90603#获取企业永久授权码
type ReqGetPermanentCodeService struct {
	// AuthCode <a href="#10974">临时授权码</a>，会在授权成功时附加在redirect_uri中跳转回第三方服务商网站，或通过<a href="#14951">授权成功通知</a>回调推送给服务商。长度为64至512个字节，必填
	AuthCode string `json:"auth_code"`
}

var _ bodyer = ReqGetPermanentCodeService{}

func (x ReqGetPermanentCodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetPermanentCodeService 获取企业永久授权码响应
// 文档：https://developer.work.weixin.qq.com/document/path/90603#获取企业永久授权码
type (
	RespGetPermanentCodeService struct {
		AccessToken  string       `json:"access_token"`
		AuthCorpInfo AuthCorpInfo `json:"auth_corp_info"`
		AuthInfo     AuthInfo     `json:"auth_info"`
		AuthUserInfo struct {
			Avatar     string `json:"avatar"`
			Name       string `json:"name"`
			OpenUserid string `json:"open_userid"`
			Userid     string `json:"userid"`
		} `json:"auth_user_info"`
		DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
		CommonResp
		ExpiresIn        int    `json:"expires_in"`
		PermanentCode    string `json:"permanent_code"`
		RegisterCodeInfo struct {
			RegisterCode string `json:"register_code"`
			State        string `json:"state"`
			TemplateID   string `json:"template_id"`
		} `json:"register_code_info"`
		// 收银台获取企业永久授权码额外返回字段：https://developer.work.weixin.qq.com/document/path/91911
		EditionInfo EditionInfo `json:"edition_info"`
	}

	AuthCorpInfo struct {
		CorpFullName      string `json:"corp_full_name"`
		CorpIndustry      string `json:"corp_industry"`
		CorpName          string `json:"corp_name"`
		CorpScale         string `json:"corp_scale"`
		CorpSquareLogoURL string `json:"corp_square_logo_url"`
		CorpRoundLogoURL  string `json:"corp_round_logo_url"`
		CorpSubIndustry   string `json:"corp_sub_industry"`
		CorpType          string `json:"corp_type"`
		CorpUserMax       int    `json:"corp_user_max"`
		CorpWxqrcode      string `json:"corp_wxqrcode"`
		Corpid            string `json:"corpid"`
		SubjectType       int    `json:"subject_type"`
		VerifiedEndTime   int    `json:"verified_end_time"`
		Location          string `json:"location"`
	}

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
	}

	DealerCorpInfo struct {
		CorpName string `json:"corp_name"`
		Corpid   string `json:"corpid"`
	}

	EditionInfo struct {
		Agent []EditionInfoAgent `json:"agent"`
	}

	EditionInfoAgent struct {
		Agentid               int    `json:"agentid"`
		EditionId             string `json:"edition_id"`
		EditionName           string `json:"edition_name"`
		AppStatus             int    `json:"app_status"`
		UserLimit             int    `json:"user_limit"`
		ExpiredTime           int    `json:"expired_time"`
		IsVirtualVersion      bool   `json:"is_virtual_version"`
		IsSharedFromOtherCorp bool   `json:"is_shared_from_other_corp"`
	}
)

var _ bodyer = RespGetPermanentCodeService{}

func (x RespGetPermanentCodeService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetPermanentCodeService 获取企业永久授权码
// 文档：https://developer.work.weixin.qq.com/document/path/90603#获取企业永久授权码
func (c *ApiClient) ExecGetPermanentCodeService(req ReqGetPermanentCodeService) (RespGetPermanentCodeService, error) {
	var resp RespGetPermanentCodeService
	err := c.executeWXApiPost("/cgi-bin/service/get_permanent_code", req, &resp, true)
	if err != nil {
		return RespGetPermanentCodeService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetPermanentCodeService{}, bizErr
	}
	return resp, nil
}
