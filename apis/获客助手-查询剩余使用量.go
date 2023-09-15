package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCustomerAcquisitionQuotaExternalcontact 查询剩余使用量请求
// 文档：https://developer.work.weixin.qq.com/document/path/97375#查询剩余使用量
type ReqCustomerAcquisitionQuotaExternalcontact struct{}

var _ urlValuer = ReqCustomerAcquisitionQuotaExternalcontact{}

func (x ReqCustomerAcquisitionQuotaExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespCustomerAcquisitionQuotaExternalcontact 查询剩余使用量响应
// 文档：https://developer.work.weixin.qq.com/document/path/97375#查询剩余使用量
type RespCustomerAcquisitionQuotaExternalcontact struct {
	// Balance 剩余使用量
	Balance int `json:"balance"`
	CommonResp
	// Total 历史累计使用量
	Total int `json:"total"`
	// QuotaList 过期额度
	QuotaList []struct {
		ExpireDate int `json:"expire_date"`
		Balance    int `json:"balance"`
	} `json:"quota_list"`
}

var _ bodyer = RespCustomerAcquisitionQuotaExternalcontact{}

func (x RespCustomerAcquisitionQuotaExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCustomerAcquisitionQuotaExternalcontact 查询剩余使用量
// 文档：https://developer.work.weixin.qq.com/document/path/97375#查询剩余使用量
func (c *ApiClient) ExecCustomerAcquisitionQuotaExternalcontact(req ReqCustomerAcquisitionQuotaExternalcontact) (RespCustomerAcquisitionQuotaExternalcontact, error) {
	var resp RespCustomerAcquisitionQuotaExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/customer_acquisition_quota", req, &resp, true)
	if err != nil {
		return RespCustomerAcquisitionQuotaExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCustomerAcquisitionQuotaExternalcontact{}, bizErr
	}
	return resp, nil
}
