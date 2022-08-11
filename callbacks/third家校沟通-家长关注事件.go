package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92051#家长关注事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeSchoolContactSubscribe{})
}

type ThirdChangeSchoolContactSubscribe struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
}

func (ThirdChangeSchoolContactSubscribe) GetMessageType() string {
	return "third"
}

func (ThirdChangeSchoolContactSubscribe) GetEventType() string {
	return "change_school_contact"
}

func (ThirdChangeSchoolContactSubscribe) GetChangeType() string {
	return "subscribe"
}

func (m ThirdChangeSchoolContactSubscribe) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeSchoolContactSubscribe) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeSchoolContactSubscribe
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
