package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92130#删除跟进成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeExternalContactDelFollowUser{})
}

type EventChangeExternalContactDelFollowUser struct {
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
	ExternalUserID struct {
		Text string `xml:",chardata"`
	} `xml:"ExternalUserID"`
}

func (EventChangeExternalContactDelFollowUser) GetMessageType() string {
	return "event"
}

func (EventChangeExternalContactDelFollowUser) GetEventType() string {
	return "change_external_contact"
}

func (EventChangeExternalContactDelFollowUser) GetChangeType() string {
	return "del_follow_user"
}

func (m EventChangeExternalContactDelFollowUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeExternalContactDelFollowUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeExternalContactDelFollowUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
