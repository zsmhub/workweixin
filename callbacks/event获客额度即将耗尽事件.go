package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97299#获客额度即将耗尽事件

func init() {
	//添加可解析的回调事件
	supportCallback(EventCustomerAcquisitionBalanceLow{})
}

type EventCustomerAcquisitionBalanceLow struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	FromUserName struct {
		Text string `xml:",chardata"`
	} `xml:"FromUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (EventCustomerAcquisitionBalanceLow) GetMessageType() string {
	return "event"
}

func (EventCustomerAcquisitionBalanceLow) GetEventType() string {
	return "customer_acquisition"
}

func (EventCustomerAcquisitionBalanceLow) GetChangeType() string {
	return "balance_low"
}

func (m EventCustomerAcquisitionBalanceLow) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventCustomerAcquisitionBalanceLow) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventCustomerAcquisitionBalanceLow
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
