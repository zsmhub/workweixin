package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#企业客户标签重排事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalTagShuffle{})
}

type ThirdChangeExternalTagShuffle struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (ThirdChangeExternalTagShuffle) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalTagShuffle) GetEventType() string {
	return "change_external_tag"
}

func (ThirdChangeExternalTagShuffle) GetChangeType() string {
	return "shuffle"
}

func (m ThirdChangeExternalTagShuffle) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalTagShuffle) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalTagShuffle
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
