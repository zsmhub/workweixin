package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#客户群解散事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalChatDismiss{})
}

type EventChangeExternalChatDismiss struct {
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
	ChatId struct {
		Text string `xml:",chardata"`
	} `xml:"ChatId"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (EventChangeExternalChatDismiss) GetMessageType() string {
	return "event"
}

func (EventChangeExternalChatDismiss) GetEventType() string {
	return "change_external_chat"
}

func (EventChangeExternalChatDismiss) GetChangeType() string {
	return "dismiss"
}

func (m EventChangeExternalChatDismiss) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalChatDismiss) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalChatDismiss
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
