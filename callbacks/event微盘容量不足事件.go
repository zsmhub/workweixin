package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95846#微盘容量不足事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventWedriveInsufficientCapacity{})
}

type EventWedriveInsufficientCapacity struct {
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
}

func (EventWedriveInsufficientCapacity) GetMessageType() string {
	return "event"
}

func (EventWedriveInsufficientCapacity) GetEventType() string {
	return "wedrive_insufficient_capacity"
}

func (EventWedriveInsufficientCapacity) GetChangeType() string {
	return ""
}

func (m EventWedriveInsufficientCapacity) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventWedriveInsufficientCapacity) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventWedriveInsufficientCapacity
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
