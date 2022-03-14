package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#进入应用

func init() {
	// 添加可解析的回调事件
	supportCallback(EventEnterAgent{})
}

type EventEnterAgent struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventEnterAgent) GetMessageType() string {
	return "event"
}

func (EventEnterAgent) GetEventType() string {
	return "enter_agent"
}

func (EventEnterAgent) GetChangeType() string {
	return ""
}

func (m EventEnterAgent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventEnterAgent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventEnterAgent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
