package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92051#删除学生事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeSchoolContactDeleteStudent{})
}

type ThirdChangeSchoolContactDeleteStudent struct {
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

func (ThirdChangeSchoolContactDeleteStudent) GetMessageType() string {
	return "third"
}

func (ThirdChangeSchoolContactDeleteStudent) GetEventType() string {
	return "change_school_contact"
}

func (ThirdChangeSchoolContactDeleteStudent) GetChangeType() string {
	return "delete_student"
}

func (m ThirdChangeSchoolContactDeleteStudent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeSchoolContactDeleteStudent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeSchoolContactDeleteStudent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
