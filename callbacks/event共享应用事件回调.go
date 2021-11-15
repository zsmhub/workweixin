package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90376#共享应用事件回调

func init() {
	// 添加可解析的回调事件
	supportCallback(EventShareAgentChange{})
}

type EventShareAgentChange struct {
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

func (EventShareAgentChange) GetMessageType() string {
	return "event"
}

func (EventShareAgentChange) GetEventType() string {
	return "share_agent_change"
}

func (EventShareAgentChange) GetChangeType() string {
	return ""
}

func (m EventShareAgentChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventShareAgentChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventShareAgentChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
