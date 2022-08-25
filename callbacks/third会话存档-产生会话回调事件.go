package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95039#产生会话回调事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdMsgauditNotify{})
}

type ThirdMsgauditNotify struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
}

func (ThirdMsgauditNotify) GetMessageType() string {
	return "third"
}

func (ThirdMsgauditNotify) GetEventType() string {
	return "msgaudit_notify"
}

func (ThirdMsgauditNotify) GetChangeType() string {
	return ""
}

func (m ThirdMsgauditNotify) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdMsgauditNotify) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdMsgauditNotify
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
