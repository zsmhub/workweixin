package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#点击菜单跳转链接的事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventView{})
}

type EventView struct {
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

func (EventView) GetMessageType() string {
	return "event"
}

func (EventView) GetEventType() string {
	return "view"
}

func (EventView) GetChangeType() string {
	return ""
}

func (m EventView) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventView) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventView
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
