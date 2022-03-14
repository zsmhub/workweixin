package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#外部联系人免验证添加成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactAddHalfExternalContact{})
}

type ThirdChangeExternalContactAddHalfExternalContact struct {
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

func (ThirdChangeExternalContactAddHalfExternalContact) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactAddHalfExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactAddHalfExternalContact) GetChangeType() string {
	return "add_half_external_contact"
}

func (m ThirdChangeExternalContactAddHalfExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactAddHalfExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactAddHalfExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
