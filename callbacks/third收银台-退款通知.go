package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/91932#退款通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdRefund{})
}

type ThirdRefund struct {
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
	OrderId struct {
		Text string `xml:",chardata"`
	} `xml:"OrderId"`
}

func (ThirdRefund) GetMessageType() string {
	return "third"
}

func (ThirdRefund) GetEventType() string {
	return "refund"
}

func (ThirdRefund) GetChangeType() string {
	return ""
}

func (m ThirdRefund) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdRefund) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdRefund
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
