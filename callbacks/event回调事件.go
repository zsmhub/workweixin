package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/94699#回调事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventKfMsgOrEvent{})
}

type EventKfMsgOrEvent struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
	Token struct {
		Text string `xml:",chardata"`
	} `xml:"Token"`
	OpenKfId struct {
		Text string `xml:",chardata"`
	} `xml:"OpenKfId"`
}

func (EventKfMsgOrEvent) GetMessageType() string {
	return "event"
}

func (EventKfMsgOrEvent) GetEventType() string {
	return "kf_msg_or_event"
}

func (EventKfMsgOrEvent) GetChangeType() string {
	return ""
}

func (m EventKfMsgOrEvent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventKfMsgOrEvent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventKfMsgOrEvent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
