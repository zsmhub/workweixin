package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#长期未使用应用重新启用事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventReopenInactiveAgent{})
}

type EventReopenInactiveAgent struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventReopenInactiveAgent) GetMessageType() string {
	return "event"
}

func (EventReopenInactiveAgent) GetEventType() string {
	return "reopen_inactive_agent"
}

func (EventReopenInactiveAgent) GetChangeType() string {
	return ""
}

func (m EventReopenInactiveAgent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventReopenInactiveAgent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventReopenInactiveAgent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
