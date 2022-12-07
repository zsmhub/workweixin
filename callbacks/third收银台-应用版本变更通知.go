package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/91933#应用版本变更通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdChangeEditon{})
}

type ThirdChangeEditon struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	PaidCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"PaidCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
}

func (ThirdChangeEditon) GetMessageType() string {
	return "third"
}

func (ThirdChangeEditon) GetEventType() string {
	return "change_editon"
}

func (ThirdChangeEditon) GetChangeType() string {
	return ""
}

func (m ThirdChangeEditon) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeEditon) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeEditon
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
