package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#编辑企业客户事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalContactEditExternalContact{})
}

type EventChangeExternalContactEditExternalContact struct {
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
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	ExternalUserID struct {
		Text string `xml:",chardata"`
	} `xml:"ExternalUserID"`
}

func (EventChangeExternalContactEditExternalContact) GetMessageType() string {
	return "event"
}

func (EventChangeExternalContactEditExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (EventChangeExternalContactEditExternalContact) GetChangeType() string {
	return "edit_external_contact"
}

func (m EventChangeExternalContactEditExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalContactEditExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalContactEditExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
