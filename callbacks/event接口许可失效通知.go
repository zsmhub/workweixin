package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95716#接口许可失效通知

func init() {
	// 添加可解析的回调事件
	supportCallback(EventUnlicensedNotify{})
}

type EventUnlicensedNotify struct {
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

func (EventUnlicensedNotify) GetMessageType() string {
	return "event"
}

func (EventUnlicensedNotify) GetEventType() string {
	return "unlicensed_notify"
}

func (EventUnlicensedNotify) GetChangeType() string {
	return ""
}

func (m EventUnlicensedNotify) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventUnlicensedNotify) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventUnlicensedNotify
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
