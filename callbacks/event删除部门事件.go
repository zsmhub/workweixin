package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#删除部门事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeContactDeleteParty{})
}

type EventChangeContactDeleteParty struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
}

func (EventChangeContactDeleteParty) GetMessageType() string {
	return "event"
}

func (EventChangeContactDeleteParty) GetEventType() string {
	return "change_contact"
}

func (EventChangeContactDeleteParty) GetChangeType() string {
	return "delete_party"
}

func (m EventChangeContactDeleteParty) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeContactDeleteParty) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeContactDeleteParty
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
