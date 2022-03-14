package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#编辑企业客户事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactEditExternalContact{})
}

type ThirdChangeExternalContactEditExternalContact struct {
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
}

func (ThirdChangeExternalContactEditExternalContact) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactEditExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactEditExternalContact) GetChangeType() string {
	return "edit_external_contact"
}

func (m ThirdChangeExternalContactEditExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactEditExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactEditExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
