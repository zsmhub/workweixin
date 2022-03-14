package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#更新部门事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeContactUpdateParty{})
}

type EventChangeContactUpdateParty struct {
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
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"Name"`
	ParentId struct {
		Text string `xml:",chardata"`
	} `xml:"ParentId"`
}

func (EventChangeContactUpdateParty) GetMessageType() string {
	return "event"
}

func (EventChangeContactUpdateParty) GetEventType() string {
	return "change_contact"
}

func (EventChangeContactUpdateParty) GetChangeType() string {
	return "update_party"
}

func (m EventChangeContactUpdateParty) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeContactUpdateParty) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeContactUpdateParty
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
