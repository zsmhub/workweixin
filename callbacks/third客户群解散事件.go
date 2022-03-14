package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#客户群解散事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalChatDismiss{})
}

type ThirdChangeExternalChatDismiss struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	ChatId struct {
		Text string `xml:",chardata"`
	} `xml:"ChatId"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (ThirdChangeExternalChatDismiss) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalChatDismiss) GetEventType() string {
	return "change_external_chat"
}

func (ThirdChangeExternalChatDismiss) GetChangeType() string {
	return "dismiss"
}

func (m ThirdChangeExternalChatDismiss) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalChatDismiss) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalChatDismiss
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
