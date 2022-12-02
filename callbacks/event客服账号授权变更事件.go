package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97712#客服账号授权变更事件

func init() {
	//添加可解析的回调事件
	supportCallback(EventKfAccountAuthChange{})
}

type EventKfAccountAuthChange struct {
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
	AuthAddOpenKfId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthAddOpenKfId"`
	AuthDelOpenKfId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthDelOpenKfId"`
}

func (EventKfAccountAuthChange) GetMessageType() string {
	return "event"
}

func (EventKfAccountAuthChange) GetEventType() string {
	return "kf_account_auth_change"
}

func (EventKfAccountAuthChange) GetChangeType() string {
	return ""
}

func (m EventKfAccountAuthChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventKfAccountAuthChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventKfAccountAuthChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
