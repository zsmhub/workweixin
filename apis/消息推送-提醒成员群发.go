package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqRemindGroupmsgSendExternalcontact 提醒成员群发请求
// 文档：https://developer.work.weixin.qq.com/document/path/97610#提醒成员群发
type ReqRemindGroupmsgSendExternalcontact struct {
	// Msgid 群发消息的id，通过获取群发记录列表接口返回，必填
	Msgid string `json:"msgid"`
}

var _ bodyer = ReqRemindGroupmsgSendExternalcontact{}

func (x ReqRemindGroupmsgSendExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespRemindGroupmsgSendExternalcontact 提醒成员群发响应
// 文档：https://developer.work.weixin.qq.com/document/path/97610#提醒成员群发
type RespRemindGroupmsgSendExternalcontact struct {
	CommonResp
}

var _ bodyer = RespRemindGroupmsgSendExternalcontact{}

func (x RespRemindGroupmsgSendExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecRemindGroupmsgSendExternalcontact 提醒成员群发
// 文档：https://developer.work.weixin.qq.com/document/path/97610#提醒成员群发
func (c *ApiClient) ExecRemindGroupmsgSendExternalcontact(req ReqRemindGroupmsgSendExternalcontact) (RespRemindGroupmsgSendExternalcontact, error) {
	var resp RespRemindGroupmsgSendExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/remind_groupmsg_send", req, &resp, true)
	if err != nil {
		return RespRemindGroupmsgSendExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespRemindGroupmsgSendExternalcontact{}, bizErr
	}
	return resp, nil
}
