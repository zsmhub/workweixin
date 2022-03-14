package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90375#语音消息

func init() {
	// 添加可解析的回调事件
	supportCallback(Voice{})
}

type Voice struct {
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
	MediaId struct {
		Text string `xml:",chardata"`
	} `xml:"MediaId"`
	Format struct {
		Text string `xml:",chardata"`
	} `xml:"Format"`
	MsgId struct {
		Text string `xml:",chardata"`
	} `xml:"MsgId"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (Voice) GetMessageType() string {
	return "voice"
}

func (Voice) GetEventType() string {
	return ""
}

func (Voice) GetChangeType() string {
	return ""
}

func (m Voice) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (Voice) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp Voice
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
