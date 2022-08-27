package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuserinfo3RdService 获取访问用户身份请求
// 文档：https://developer.work.weixin.qq.com/document/path/91121#获取访问用户身份
type ReqGetuserinfo3RdService struct {
	// Code 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。，必填
	Code string `json:"code"`
}

var _ urlValuer = ReqGetuserinfo3RdService{}

func (x ReqGetuserinfo3RdService) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetuserinfo3RdService 获取访问用户身份响应
// 文档：https://developer.work.weixin.qq.com/document/path/91121#获取访问用户身份
type RespGetuserinfo3RdService struct {
	CommonResp
	CorpID     string `json:"CorpId"`
	UserID     string `json:"UserId"`
	DeviceID   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserID string `json:"open_userid"`
}

var _ bodyer = RespGetuserinfo3RdService{}

func (x RespGetuserinfo3RdService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetuserinfo3RdService 获取访问用户身份
// 文档：https://developer.work.weixin.qq.com/document/path/91121#获取访问用户身份
func (c *ApiClient) ExecGetuserinfo3RdService(req ReqGetuserinfo3RdService) (RespGetuserinfo3RdService, error) {
	var resp RespGetuserinfo3RdService
	err := c.executeWXApiGet("/cgi-bin/service/getuserinfo3rd", req, &resp, true)
	if err != nil {
		return RespGetuserinfo3RdService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuserinfo3RdService{}, bizErr
	}
	return resp, nil
}
