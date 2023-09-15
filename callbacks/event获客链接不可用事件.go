package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97299#获客链接不可用事件

func init() {
	//添加可解析的回调事件
	supportCallback(EventCustomerAcquisitionLinkUnavailable{})
}

type EventCustomerAcquisitionLinkUnavailable struct {
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
	LinkId struct {
		Text string `xml:",chardata"`
	} `xml:"LinkId"`
}

func (EventCustomerAcquisitionLinkUnavailable) GetMessageType() string {
	return "event"
}

func (EventCustomerAcquisitionLinkUnavailable) GetEventType() string {
	return "customer_acquisition"
}

func (EventCustomerAcquisitionLinkUnavailable) GetChangeType() string {
	return "link_unavailable"
}

func (m EventCustomerAcquisitionLinkUnavailable) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventCustomerAcquisitionLinkUnavailable) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventCustomerAcquisitionLinkUnavailable
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
