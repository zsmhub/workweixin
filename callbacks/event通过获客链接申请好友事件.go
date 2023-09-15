package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97299#通过获客链接申请好友事件

func init() {
	//添加可解析的回调事件
	supportCallback(EventCustomerAcquisitionFriendRequest{})
}

type EventCustomerAcquisitionFriendRequest struct {
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
	LinkId struct {
		Text string `xml:",chardata"`
	} `xml:"LinkId"`
	State struct {
		Text string `xml:",chardata"`
	} `xml:"State"`
}

func (EventCustomerAcquisitionFriendRequest) GetMessageType() string {
	return "event"
}

func (EventCustomerAcquisitionFriendRequest) GetEventType() string {
	return "customer_acquisition"
}

func (EventCustomerAcquisitionFriendRequest) GetChangeType() string {
	return "friend_request"
}

func (m EventCustomerAcquisitionFriendRequest) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventCustomerAcquisitionFriendRequest) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventCustomerAcquisitionFriendRequest
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
