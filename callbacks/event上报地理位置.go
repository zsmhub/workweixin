package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#上报地理位置

func init() {
	// 添加可解析的回调事件
	supportCallback(EventLOCATION{})
}

type EventLOCATION struct {
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
	Latitude struct {
		Text string `xml:",chardata"`
	} `xml:"Latitude"`
	Longitude struct {
		Text string `xml:",chardata"`
	} `xml:"Longitude"`
	Precision struct {
		Text string `xml:",chardata"`
	} `xml:"Precision"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	AppType struct {
		Text string `xml:",chardata"`
	} `xml:"AppType"`
}

func (EventLOCATION) GetMessageType() string {
	return "event"
}

func (EventLOCATION) GetEventType() string {
	return "LOCATION"
}

func (EventLOCATION) GetChangeType() string {
	return ""
}

func (m EventLOCATION) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventLOCATION) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventLOCATION
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
