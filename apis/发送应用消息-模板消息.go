package apis

import (
	"bytes"
	"encoding/json"
)

// 发送应用消息-模板消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

type (
	SentTemplateMsg struct {
		TemplateId  string                       `json:"template_id"`  // 模板ID。第三方管理端创建模板后获得。对于正式授权的应用，需要审批通过后才可使用。最长64字节
		Url         string                       `json:"url"`          // 点击模板消息后的跳转链接。最长2048字节。注意，url必须带协议头 "http://" 或 "https://" 。url和miniprogram 至少要填一个,都填时优先miniprogram。
		Miniprogram SentTemplateMsgMiniprogram   `json:"miniprogram"`  // 点击后需要跳转的小程序，miniprogram和url至少要填一个，都填时优先miniprogram。
		ContentItem []SentTemplateMsgContentItem `json:"content_item"` // 消息内容键值对，允许个数范围：1~5，实际由申请的模板样式决定
	}

	SentTemplateMsgMiniprogram struct {
		Appid    string `json:"appid"`    // 在miniprogram节点中该字段必填，小程序appid，必须是与当前应用关联的小程序
		Pagepath string `json:"pagepath"` // 在miniprogram节点中该字段必填，表示点击消息卡片后的小程序页面，仅限本小程序内的页面。
	}
	SentTemplateMsgContentItem struct {
		Key   string `json:"key"`   // 1~20个utf8字符。注意，必须与template_id对应模板匹配
		Value string `json:"value"` // 1~40个utf8字符
	}
)

type ReqSentMessageCardTemplateMsg struct {
	ToUser             string          `json:"touser,omitempty"`          // 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToParty            string          `json:"toparty,omitempty"`         // 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag              string          `json:"totag,omitempty"`           // 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	MsgType            string          `json:"msgtype"`                   // 消息类型，如text/image/file
	AgentId            int             `json:"agentid"`                   // 企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
	EnableIdTrans      int             `json:"enable_id_trans,omitempty"` // 表示是否开启id转译，0表示否，1表示是，默认0
	SelectedTicketList []string        `json:"selected_ticket_list"`      // 选人sdk或者选人jsapi返回的ticket列表，列表不超过10个。接收者不包含selected_tikcet的操作者，若要发送给操作者，可将操作者填到touser字段。
	TemplateMsg        SentTemplateMsg `json:"template_msg"`
	OnlyUnauth         bool            `json:"only_unauth"` // 仅向selected_ticket_list中未授权的用户发送模板消息，仅当selected_ticket_list存在时该字段生效。如果该字段为true，则自动忽略touser，toparty，totag
}

var _ bodyer = ReqSentMessageCardTemplateMsg{}

func (x ReqSentMessageCardTemplateMsg) intoBody() ([]byte, error) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(x)
	if err != nil {
		panic(err)
	}
	return byteBuf.Bytes(), nil
}

// 发送应用消息-模板消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义
func (c *ApiClient) ExecSentMessageCardTemplateMsg(req ReqSentMessageCardTemplateMsg) (RespSentMessageCard, error) {
	var resp RespSentMessageCard
	err := c.executeWXApiPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return RespSentMessageCard{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSentMessageCard{}, bizErr
	}
	return resp, nil
}
