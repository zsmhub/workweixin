package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/91930#改单通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdChangeOrder{})
}

type ThirdChangeOrder struct {
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
	OldOrderId struct {
		Text string `xml:",chardata"`
	} `xml:"OldOrderId"`
	NewOrderId struct {
		Text string `xml:",chardata"`
	} `xml:"NewOrderId"`
}

func (ThirdChangeOrder) GetMessageType() string {
	return "third"
}

func (ThirdChangeOrder) GetEventType() string {
	return "change_order"
}

func (ThirdChangeOrder) GetChangeType() string {
	return ""
}

func (m ThirdChangeOrder) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeOrder) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeOrder
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
