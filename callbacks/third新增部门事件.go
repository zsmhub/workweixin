package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92655#新增部门事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeContactCreateParty{})
}

type ThirdChangeContactCreateParty struct {
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
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"Name"`
	ParentId struct {
		Text string `xml:",chardata"`
	} `xml:"ParentId"`
	Order struct {
		Text string `xml:",chardata"`
	} `xml:"Order"`
}

func (ThirdChangeContactCreateParty) GetMessageType() string {
	return "third"
}

func (ThirdChangeContactCreateParty) GetEventType() string {
	return "change_contact"
}

func (ThirdChangeContactCreateParty) GetChangeType() string {
	return "create_party"
}

func (m ThirdChangeContactCreateParty) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeContactCreateParty) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeContactCreateParty
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
