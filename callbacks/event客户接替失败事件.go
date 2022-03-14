package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#客户接替失败事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalContactTransferFail{})
}

type EventChangeExternalContactTransferFail struct {
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
	FailReason struct {
		Text string `xml:",chardata"`
	} `xml:"FailReason"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	ExternalUserID struct {
		Text string `xml:",chardata"`
	} `xml:"ExternalUserID"`
}

func (EventChangeExternalContactTransferFail) GetMessageType() string {
	return "event"
}

func (EventChangeExternalContactTransferFail) GetEventType() string {
	return "change_external_contact"
}

func (EventChangeExternalContactTransferFail) GetChangeType() string {
	return "transfer_fail"
}

func (m EventChangeExternalContactTransferFail) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalContactTransferFail) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalContactTransferFail
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
