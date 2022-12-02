package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97495#应用邮箱接收邮件事件

func init() {
	//添加可解析的回调事件
	supportCallback(EventAppEmailChangeReceiveEmail{})
}

type EventAppEmailChangeReceiveEmail struct {
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
	Amount struct {
		Text string `xml:",chardata"`
	} `xml:"Amount"`
}

func (EventAppEmailChangeReceiveEmail) GetMessageType() string {
	return "event"
}

func (EventAppEmailChangeReceiveEmail) GetEventType() string {
	return "app_email_change"
}

func (EventAppEmailChangeReceiveEmail) GetChangeType() string {
	return "receive_email"
}

func (m EventAppEmailChangeReceiveEmail) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventAppEmailChangeReceiveEmail) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventAppEmailChangeReceiveEmail
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
