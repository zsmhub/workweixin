package apis

import (
	"encoding/json"
)

// ReqUpdateTemplateCardMessage 接口定义请求
// 文档：https://developer.work.weixin.qq.com/document/path/94945#更新按钮为不可点击状态

type (
	ReqUpdateTemplateCardMessage struct {
		Userids      []string                           `json:"userids"`       // 企业的成员ID列表（最多支持1000个）
		Partyids     []int                              `json:"partyids"`      // 企业的部门ID列表（最多支持100个）
		Tagids       []int                              `json:"tagids"`        // 企业的标签ID列表（最多支持100个）
		Atall        int                                `json:"atall"`         // 更新整个任务接收人员
		Agentid      int                                `json:"agentid"`       // 应用的agentid
		ResponseCode string                             `json:"response_code"` // 更新卡片所需要消费的code，可通过发消息接口和回调接口返回值获取，一个code只能调用一次该接口，且只能在24小时内调用
		Button       ReqUpdateTemplateCardMessageButton `json:"button"`
	}

	ReqUpdateTemplateCardMessageButton struct {
		ReplaceName string `json:"replace_name"` // 需要更新的按钮的文案
	}
)

var _ bodyer = ReqUpdateTemplateCardMessage{}

func (x ReqUpdateTemplateCardMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateTemplateCardMessage 接口定义响应
// 文档：https://developer.work.weixin.qq.com/document/path/94945#更新按钮为不可点击状态
type RespUpdateTemplateCardMessage struct {
	CommonResp
}

var _ bodyer = RespUpdateTemplateCardMessage{}

func (x RespUpdateTemplateCardMessage) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文档：https://developer.work.weixin.qq.com/document/path/94945#更新按钮为不可点击状态
func (c *ApiClient) ExecUpdateTemplateCardMessage(req ReqUpdateTemplateCardMessage) (RespUpdateTemplateCardMessage, error) {
	var resp RespUpdateTemplateCardMessage
	err := c.executeWXApiPost("/cgi-bin/message/update_template_card", req, &resp, true)
	if err != nil {
		return RespUpdateTemplateCardMessage{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateTemplateCardMessage{}, bizErr
	}
	return resp, nil
}
