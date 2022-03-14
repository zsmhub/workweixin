package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#标签成员变更事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeContactUpdateTag{})
}

type EventChangeContactUpdateTag struct {
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
	TagId struct {
		Text string `xml:",chardata"`
	} `xml:"TagId"`
	AddUserItems struct {
		Text string `xml:",chardata"`
	} `xml:"AddUserItems"`
	DelUserItems struct {
		Text string `xml:",chardata"`
	} `xml:"DelUserItems"`
	AddPartyItems struct {
		Text string `xml:",chardata"`
	} `xml:"AddPartyItems"`
	DelPartyItems struct {
		Text string `xml:",chardata"`
	} `xml:"DelPartyItems"`
}

func (EventChangeContactUpdateTag) GetMessageType() string {
	return "event"
}

func (EventChangeContactUpdateTag) GetEventType() string {
	return "change_contact"
}

func (EventChangeContactUpdateTag) GetChangeType() string {
	return "update_tag"
}

func (m EventChangeContactUpdateTag) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeContactUpdateTag) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeContactUpdateTag
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
