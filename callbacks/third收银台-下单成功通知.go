package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/91929#下单成功通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdOpenOrder{})
}

type ThirdOpenOrder struct {
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
	OperatorId struct {
		Text string `xml:",chardata"`
	} `xml:"OperatorId"`
}

func (ThirdOpenOrder) GetMessageType() string {
	return "third"
}

func (ThirdOpenOrder) GetEventType() string {
	return "open_order"
}

func (ThirdOpenOrder) GetChangeType() string {
	return ""
}

func (m ThirdOpenOrder) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdOpenOrder) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdOpenOrder
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
