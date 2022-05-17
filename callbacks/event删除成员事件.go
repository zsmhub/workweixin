package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90970#删除成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeContactDeleteUser{})
}

type EventChangeContactDeleteUser struct {
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
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
}

func (EventChangeContactDeleteUser) GetMessageType() string {
	return "event"
}

func (EventChangeContactDeleteUser) GetEventType() string {
	return "change_contact"
}

func (EventChangeContactDeleteUser) GetChangeType() string {
	return "delete_user"
}

func (m EventChangeContactDeleteUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeContactDeleteUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeContactDeleteUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
