package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://open.work.weixin.qq.com/api/doc/90000/90135/92130#企业客户标签创建事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalTagCreate{})
}

// XML was generated 2021-09-15 18:04:51 by chenjianlin on JZTech-chenjianlin.lan.
type EventChangeExternalTagCreate struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
	TagType struct {
		Text string `xml:",chardata"`
	} `xml:"TagType"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (EventChangeExternalTagCreate) GetMessageType() string {
	return "event"
}

func (EventChangeExternalTagCreate) GetEventType() string {
	return "change_external_tag"
}

func (EventChangeExternalTagCreate) GetChangeType() string {
	return "create"
}

func (m EventChangeExternalTagCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalTagCreate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalTagCreate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
