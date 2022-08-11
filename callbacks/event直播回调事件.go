package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/94308#直播回调事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventLivingStatusChange{})
}

type EventLivingStatusChange struct {
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
	LivingId struct {
		Text string `xml:",chardata"`
	} `xml:"LivingId"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"Status"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventLivingStatusChange) GetMessageType() string {
	return "event"
}

func (EventLivingStatusChange) GetEventType() string {
	return "living_status_change"
}

func (EventLivingStatusChange) GetChangeType() string {
	return ""
}

func (m EventLivingStatusChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventLivingStatusChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventLivingStatusChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
