package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCancelGroupmsgSendExternalcontact 停止企业群发请求
// 文档：https://developer.work.weixin.qq.com/document/path/97611#停止企业群发
type ReqCancelGroupmsgSendExternalcontact struct {
	// Msgid 群发消息的id，通过<a href="#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8" rel="nofollow">获取群发记录列表</a>接口返回，必填
	Msgid string `json:"msgid"`
}

var _ bodyer = ReqCancelGroupmsgSendExternalcontact{}

func (x ReqCancelGroupmsgSendExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCancelGroupmsgSendExternalcontact 停止企业群发响应
// 文档：https://developer.work.weixin.qq.com/document/path/97611#停止企业群发
type RespCancelGroupmsgSendExternalcontact struct {
	CommonResp
}

var _ bodyer = RespCancelGroupmsgSendExternalcontact{}

func (x RespCancelGroupmsgSendExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCancelGroupmsgSendExternalcontact 停止企业群发
// 文档：https://developer.work.weixin.qq.com/document/path/97611#停止企业群发
func (c *ApiClient) ExecCancelGroupmsgSendExternalcontact(req ReqCancelGroupmsgSendExternalcontact) (RespCancelGroupmsgSendExternalcontact, error) {
	var resp RespCancelGroupmsgSendExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/cancel_groupmsg_send", req, &resp, true)
	if err != nil {
		return RespCancelGroupmsgSendExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCancelGroupmsgSendExternalcontact{}, bizErr
	}
	return resp, nil
}
