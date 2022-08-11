package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/93704#修改日程事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventModifySchedule{})
}

type EventModifySchedule struct {
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
	CalId struct {
		Text string `xml:",chardata"`
	} `xml:"CalId"`
	ScheduleId struct {
		Text string `xml:",chardata"`
	} `xml:"ScheduleId"`
}

func (EventModifySchedule) GetMessageType() string {
	return "event"
}

func (EventModifySchedule) GetEventType() string {
	return "modify_schedule"
}

func (EventModifySchedule) GetChangeType() string {
	return ""
}

func (m EventModifySchedule) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventModifySchedule) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventModifySchedule
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
