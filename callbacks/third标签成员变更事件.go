package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92656#标签成员变更事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeContactUpdateTag{})
}

type ThirdChangeContactUpdateTag struct {
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
	TagId struct {
		Text string `xml:",chardata"`
	} `xml:"TagId"`
	AddUserItems struct {
		Text string `xml:",chardata"`
	} `xml:"AddUserItems"`
	DelUserItems struct {
		Text string `xml:",chardata"`
	} `xml:"DelUserItems"`
	AddPartyItems struct {
		Text string `xml:",chardata"`
	} `xml:"AddPartyItems"`
	DelPartyItems struct {
		Text string `xml:",chardata"`
	} `xml:"DelPartyItems"`
}

func (ThirdChangeContactUpdateTag) GetMessageType() string {
	return "third"
}

func (ThirdChangeContactUpdateTag) GetEventType() string {
	return "change_contact"
}

func (ThirdChangeContactUpdateTag) GetChangeType() string {
	return "update_tag"
}

func (m ThirdChangeContactUpdateTag) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeContactUpdateTag) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeContactUpdateTag
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
