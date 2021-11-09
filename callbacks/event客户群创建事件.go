package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://open.work.weixin.qq.com/api/doc/90000/90135/92130#客户群创建事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalChatCreate{})
}

// XML was generated 2021-09-15 18:04:51 by chenjianlin on JZTech-chenjianlin.lan.
type EventChangeExternalChatCreate struct {
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

func (EventChangeExternalChatCreate) GetMessageType() string {
	return "event"
}

func (EventChangeExternalChatCreate) GetEventType() string {
	return "change_external_chat"
}

func (EventChangeExternalChatCreate) GetChangeType() string {
	return "create"
}

func (m EventChangeExternalChatCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalChatCreate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalChatCreate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
