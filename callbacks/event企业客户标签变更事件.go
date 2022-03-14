package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#企业客户标签变更事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalTagUpdate{})
}

type EventChangeExternalTagUpdate struct {
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

func (EventChangeExternalTagUpdate) GetMessageType() string {
	return "event"
}

func (EventChangeExternalTagUpdate) GetEventType() string {
	return "change_external_tag"
}

func (EventChangeExternalTagUpdate) GetChangeType() string {
	return "update"
}

func (m EventChangeExternalTagUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalTagUpdate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalTagUpdate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
