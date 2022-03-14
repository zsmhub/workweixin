package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#客户群变更事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalChatUpdate{})
}

type EventChangeExternalChatUpdate struct {
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
	UpdateDetail struct {
		Text string `xml:",chardata"`
	} `xml:"UpdateDetail"`
	JoinScene struct {
		Text string `xml:",chardata"`
	} `xml:"JoinScene"`
	QuitScene struct {
		Text string `xml:",chardata"`
	} `xml:"QuitScene"`
	MemChangeCnt struct {
		Text string `xml:",chardata"`
	} `xml:"MemChangeCnt"`
}

func (EventChangeExternalChatUpdate) GetMessageType() string {
	return "event"
}

func (EventChangeExternalChatUpdate) GetEventType() string {
	return "change_external_chat"
}

func (EventChangeExternalChatUpdate) GetChangeType() string {
	return "update"
}

func (m EventChangeExternalChatUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalChatUpdate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalChatUpdate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
