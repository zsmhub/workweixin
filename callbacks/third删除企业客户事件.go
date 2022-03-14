package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#删除企业客户事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactDelExternalContact{})
}

type ThirdChangeExternalContactDelExternalContact struct {
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
	Source struct {
		Text string `xml:",chardata"`
	} `xml:"Source"`
}

func (ThirdChangeExternalContactDelExternalContact) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactDelExternalContact) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactDelExternalContact) GetChangeType() string {
	return "del_external_contact"
}

func (m ThirdChangeExternalContactDelExternalContact) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactDelExternalContact) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactDelExternalContact
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
