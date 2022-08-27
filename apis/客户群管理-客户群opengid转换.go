package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqOpengidToChatidExternalcontact 客户群opengid转换请求
// 文档：https://developer.work.weixin.qq.com/document/path/94828#客户群opengid转换
type ReqOpengidToChatidExternalcontact struct {
	// Opengid 小程序在微信获取到的群ID，参见<a href="https://developers.weixin.qq.com/miniprogram/dev/api/open-api/group/wx.getGroupEnterInfo.html" rel="nofollow">wx.getGroupEnterInfo</a>，必填
	Opengid string `json:"opengid"`
}

var _ bodyer = ReqOpengidToChatidExternalcontact{}

func (x ReqOpengidToChatidExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespOpengidToChatidExternalcontact 客户群opengid转换响应
// 文档：https://developer.work.weixin.qq.com/document/path/94828#客户群opengid转换
type RespOpengidToChatidExternalcontact struct {
	CommonResp
	ChatId string `json:"chat_id"`
}

var _ bodyer = RespOpengidToChatidExternalcontact{}

func (x RespOpengidToChatidExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecOpengidToChatidExternalcontact 客户群opengid转换
// 文档：https://developer.work.weixin.qq.com/document/path/94828#客户群opengid转换
func (c *ApiClient) ExecOpengidToChatidExternalcontact(req ReqOpengidToChatidExternalcontact) (RespOpengidToChatidExternalcontact, error) {
	var resp RespOpengidToChatidExternalcontact
	err := c.executeWXApiPost(":/cgi-bin/externalcontact/opengid_to_chatid", req, &resp, true)
	if err != nil {
		return RespOpengidToChatidExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOpengidToChatidExternalcontact{}, bizErr
	}
	return resp, nil
}
