package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95039#产生会话回调事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventMsgauditNotify{})
}

type EventMsgauditNotify struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
}

func (EventMsgauditNotify) GetMessageType() string {
	return "event"
}

func (EventMsgauditNotify) GetEventType() string {
	return "msgaudit_notify"
}

func (EventMsgauditNotify) GetChangeType() string {
	return ""
}

func (m EventMsgauditNotify) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventMsgauditNotify) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventMsgauditNotify
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
