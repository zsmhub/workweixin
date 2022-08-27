package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// https://developer.work.weixin.qq.com/document/path/91039
type ReqGetCustomizedCorpTokenService struct {
	// AuthCorpid 授权方corpid，必填
	Corpid string `json:"corpid"`
	// PermanentCode 永久授权码，通过get_permanent_code获取，必填
	Corpsecret string `json:"corpsecret"`
}

var _ urlValuer = ReqGetCustomizedCorpTokenService{}

func (x ReqGetCustomizedCorpTokenService) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// https://developer.work.weixin.qq.com/document/path/91039
type RespGetCustomizedCorpTokenService struct {
	CommonResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var _ bodyer = RespGetCustomizedCorpTokenService{}

func (x RespGetCustomizedCorpTokenService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文档：https://developer.work.weixin.qq.com/document/path/91039
func (c *ApiClient) ExecGetCustomizedCorpTokenService(req ReqGetCustomizedCorpTokenService) (RespGetCustomizedCorpTokenService, error) {
	var resp RespGetCustomizedCorpTokenService
	err := c.executeWXApiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return RespGetCustomizedCorpTokenService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetCustomizedCorpTokenService{}, bizErr
	}
	return resp, nil
}
