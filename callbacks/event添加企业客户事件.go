package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#添加企业客户事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalContactAddExternalContact{})
}

type EventChangeExternalContactAddExternalContact struct {
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
	State struct {
		Text string `xml:",chardata"`
	} `xml:"State"`
	WelcomeCode struct {
		Text string `xml:",chardata"`
	} `xml:"WelcomeCode"`
}

func (EventChangeExternalContactAddExternalContact) GetMessageType() string {
	return "event"
}

func (EventChangeExternalContactAddExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (EventChangeExternalContactAddExternalContact) GetChangeType() string {
	return "add_external_contact"
}

func (m EventChangeExternalContactAddExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalContactAddExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalContactAddExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
