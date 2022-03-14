package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90375#位置消息

func init() {
	// 添加可解析的回调事件
	supportCallback(Location{})
}

type Location struct {
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
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"Location_X"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"Location_Y"`
	Scale struct {
		Text string `xml:",chardata"`
	} `xml:"Scale"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"Label"`
	MsgId struct {
		Text string `xml:",chardata"`
	} `xml:"MsgId"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	AppType struct {
		Text string `xml:",chardata"`
	} `xml:"AppType"`
}

func (Location) GetMessageType() string {
	return "location"
}

func (Location) GetEventType() string {
	return ""
}

func (Location) GetChangeType() string {
	return ""
}

func (m Location) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (Location) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp Location
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
