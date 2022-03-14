package callbacks

import "encoding/xml"

// 文档: https://developer.work.weixin.qq.com/document/path/90376#成员关注及取消关注事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventUnsubscribe{})
}

type EventUnsubscribe struct {
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

func (EventUnsubscribe) GetMessageType() string {
	return "event"
}

func (EventUnsubscribe) GetEventType() string {
	return "unsubscribe"
}

func (EventUnsubscribe) GetChangeType() string {
	return ""
}

func (m EventUnsubscribe) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventUnsubscribe) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventUnsubscribe
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
