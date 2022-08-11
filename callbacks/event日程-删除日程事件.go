package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/93704#删除日程事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventDeleteSchedule{})
}

type EventDeleteSchedule struct {
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

func (EventDeleteSchedule) GetMessageType() string {
	return "event"
}

func (EventDeleteSchedule) GetEventType() string {
	return "delete_schedule"
}

func (EventDeleteSchedule) GetChangeType() string {
	return ""
}

func (m EventDeleteSchedule) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventDeleteSchedule) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventDeleteSchedule
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
