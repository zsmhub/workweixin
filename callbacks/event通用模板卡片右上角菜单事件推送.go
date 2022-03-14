package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#通用模板卡片右上角菜单事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventTemplateCardMenuEvent{})
}

type EventTemplateCardMenuEvent struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	FromUserName struct {
		Text string `xml:",chardata"`
	} `xml:"FromUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
	EventKey struct {
		Text string `xml:",chardata"`
	} `xml:"EventKey"`
	TaskId struct {
		Text string `xml:",chardata"`
	} `xml:"TaskId"`
	CardType struct {
		Text string `xml:",chardata"`
	} `xml:"CardType"`
	ResponseCode struct {
		Text string `xml:",chardata"`
	} `xml:"ResponseCode"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventTemplateCardMenuEvent) GetMessageType() string {
	return "event"
}

func (EventTemplateCardMenuEvent) GetEventType() string {
	return "template_card_menu_event"
}

func (EventTemplateCardMenuEvent) GetChangeType() string {
	return ""
}

func (m EventTemplateCardMenuEvent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventTemplateCardMenuEvent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventTemplateCardMenuEvent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
