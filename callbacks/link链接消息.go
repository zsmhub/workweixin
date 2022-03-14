package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90375#链接消息

func init() {
	// 添加可解析的回调事件
	supportCallback(Link{})
}

type Link struct {
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
	Title struct {
		Text string `xml:",chardata"`
	} `xml:"Title"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"Description"`
	URL struct {
		Text string `xml:",chardata"`
	} `xml:"Url"`
	PicUrl struct {
		Text string `xml:",chardata"`
	} `xml:"PicUrl"`
	MsgId struct {
		Text string `xml:",chardata"`
	} `xml:"MsgId"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (Link) GetMessageType() string {
	return "link"
}

func (Link) GetEventType() string {
	return ""
}

func (Link) GetChangeType() string {
	return ""
}

func (m Link) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (Link) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp Link
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
