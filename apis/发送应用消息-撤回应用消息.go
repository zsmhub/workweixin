package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqRecallMessage 撤回应用消息请求
// 文档：https://developer.work.weixin.qq.com/document/path/94947#撤回应用消息
type ReqRecallMessage struct {
	// Msgid 消息ID。从应用发送消息接口处获得。，必填
	Msgid string `json:"msgid"`
}

var _ bodyer = ReqRecallMessage{}

func (x ReqRecallMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespRecallMessage 撤回应用消息响应
// 文档：https://developer.work.weixin.qq.com/document/path/94947#撤回应用消息
type RespRecallMessage struct {
	CommonResp
}

var _ bodyer = RespRecallMessage{}

func (x RespRecallMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecRecallMessage 撤回应用消息
// 文档：https://developer.work.weixin.qq.com/document/path/94947#撤回应用消息
func (c *ApiClient) ExecRecallMessage(req ReqRecallMessage) (RespRecallMessage, error) {
	var resp RespRecallMessage
	err := c.executeWXApiPost("/cgi-bin/message/recall", req, &resp, true)
	if err != nil {
		return RespRecallMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespRecallMessage{}, bizErr
	}
	return resp, nil
}
