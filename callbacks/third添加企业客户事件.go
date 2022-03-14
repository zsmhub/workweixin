package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#添加企业客户事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactAddExternalContact{})
}

type ThirdChangeExternalContactAddExternalContact struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
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

func (ThirdChangeExternalContactAddExternalContact) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactAddExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactAddExternalContact) GetChangeType() string {
	return "add_external_contact"
}

func (m ThirdChangeExternalContactAddExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactAddExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactAddExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
