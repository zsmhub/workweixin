package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#扫码推事件的事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventScancodePush{})
}

type EventScancodePush struct {
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
	ScanCodeInfo struct {
		Text     string `xml:",chardata"`
		ScanType struct {
			Text string `xml:",chardata"`
		} `xml:"ScanType"`
		ScanResult struct {
			Text string `xml:",chardata"`
		} `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventScancodePush) GetMessageType() string {
	return "event"
}

func (EventScancodePush) GetEventType() string {
	return "scancode_push"
}

func (EventScancodePush) GetChangeType() string {
	return ""
}

func (m EventScancodePush) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventScancodePush) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventScancodePush
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
