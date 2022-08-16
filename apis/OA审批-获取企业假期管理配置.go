package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetcorpconfVacation 获取企业假期管理配置请求
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置

type ReqGetcorpconfVacation struct {
}

var _ urlValuer = ReqGetcorpconfVacation{}

func (x ReqGetcorpconfVacation) intoURLValues() url.Values {
	var ret url.Values = make(map[string][]string)

	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetcorpconfVacation 获取企业假期管理配置响应
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置

type RespGetcorpconfVacation struct {
	CommonResp
	Lists []struct {
		DurationType   int    `json:"duration_type"`
		ID             int    `json:"id"`
		Name           string `json:"name"`
		PerdayDuration int    `json:"perday_duration"`
		QuotaAttr      struct {
			AutoresetDuration int `json:"autoreset_duration"`
			AutoresetTime     int `json:"autoreset_time"`
			Type              int `json:"type"`
		} `json:"quota_attr"`
		TimeAttr int `json:"time_attr"`
	} `json:"lists"`
}

var _ bodyer = RespGetcorpconfVacation{}

func (x RespGetcorpconfVacation) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// execGetcorpconfVacation 获取企业假期管理配置
// 文档：https://developer.work.weixin.qq.com/document/path/94211#获取企业假期管理配置
func (c *ApiClient) ExecGetcorpconfVacation(req ReqGetcorpconfVacation) (RespGetcorpconfVacation, error) {
	var resp RespGetcorpconfVacation
	err := c.executeWXApiGet("/cgi-bin/oa/vacation/getcorpconf", req, &resp, true)
	if err != nil {
		return RespGetcorpconfVacation{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetcorpconfVacation{}, bizErr
	}

	return resp, nil
}
