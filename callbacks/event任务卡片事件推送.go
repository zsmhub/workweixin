package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#任务卡片事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventTaskcardClick{})
}

type EventTaskcardClick struct {
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
	EventKey struct {
		Text string `xml:",chardata"`
	} `xml:"EventKey"`
	TaskId struct {
		Text string `xml:",chardata"`
	} `xml:"TaskId"`
	AgentId struct {
		Text string `xml:",chardata"`
	} `xml:"AgentId"`
}

func (EventTaskcardClick) GetMessageType() string {
	return "event"
}

func (EventTaskcardClick) GetEventType() string {
	return "taskcard_click"
}

func (EventTaskcardClick) GetChangeType() string {
	return ""
}

func (m EventTaskcardClick) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventTaskcardClick) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventTaskcardClick
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
