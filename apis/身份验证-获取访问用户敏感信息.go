package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetuserdetail3RdService 获取访问用户敏感信息请求
// 文档：https://developer.work.weixin.qq.com/document/path/91122#获取访问用户敏感信息
type ReqGetuserdetail3RdService struct {
	// UserTicket 成员票据，必填
	UserTicket string `json:"user_ticket"`
}

var _ bodyer = ReqGetuserdetail3RdService{}

func (x ReqGetuserdetail3RdService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetuserdetail3RdService 获取访问用户敏感信息响应
// 文档：https://developer.work.weixin.qq.com/document/path/91122#获取访问用户敏感信息
type RespGetuserdetail3RdService struct {
	Avatar string `json:"avatar"`
	Corpid string `json:"corpid"`
	CommonResp
	Gender string `json:"gender"`
	Name   string `json:"name"`
	QrCode string `json:"qr_code"`
	Userid string `json:"userid"`
}

var _ bodyer = RespGetuserdetail3RdService{}

func (x RespGetuserdetail3RdService) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetuserdetail3RdService 获取访问用户敏感信息
// 文档：https://developer.work.weixin.qq.com/document/path/91122#获取访问用户敏感信息
func (c *ApiClient) ExecGetuserdetail3RdService(req ReqGetuserdetail3RdService) (RespGetuserdetail3RdService, error) {
	var resp RespGetuserdetail3RdService
	err := c.executeWXApiPost("/cgi-bin/service/getuserdetail3rd", req, &resp, true)
	if err != nil {
		return RespGetuserdetail3RdService{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetuserdetail3RdService{}, bizErr
	}
	return resp, nil
}
