package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92005#客户同意进行聊天内容存档事件回调

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactMsgAuditApproved{})
}

type ThirdChangeExternalContactMsgAuditApproved struct {
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
	WelcomeCode struct {
		Text string `xml:",chardata"`
	} `xml:"WelcomeCode"`
}

func (ThirdChangeExternalContactMsgAuditApproved) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactMsgAuditApproved) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactMsgAuditApproved) GetChangeType() string {
	return "msg_audit_approved"
}

func (m ThirdChangeExternalContactMsgAuditApproved) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactMsgAuditApproved) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactMsgAuditApproved
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
