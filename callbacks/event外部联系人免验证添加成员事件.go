package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#外部联系人免验证添加成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalContactAddHalfExternalContact{})
}

type EventChangeExternalContactAddHalfExternalContact struct {
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

func (EventChangeExternalContactAddHalfExternalContact) GetMessageType() string {
	return "event"
}

func (EventChangeExternalContactAddHalfExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (EventChangeExternalContactAddHalfExternalContact) GetChangeType() string {
	return "add_half_external_contact"
}

func (m EventChangeExternalContactAddHalfExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalContactAddHalfExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalContactAddHalfExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
