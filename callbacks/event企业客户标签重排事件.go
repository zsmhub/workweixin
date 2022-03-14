package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#企业客户标签重排事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalTagShuffle{})
}

type EventChangeExternalTagShuffle struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
	StrategyId struct {
		Text string `xml:",chardata"`
	} `xml:"StrategyId"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (EventChangeExternalTagShuffle) GetMessageType() string {
	return "event"
}

func (EventChangeExternalTagShuffle) GetEventType() string {
	return "change_external_tag"
}

func (EventChangeExternalTagShuffle) GetChangeType() string {
	return "shuffle"
}

func (m EventChangeExternalTagShuffle) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalTagShuffle) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalTagShuffle
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
