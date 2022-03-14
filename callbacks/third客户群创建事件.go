package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#客户群创建事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalChatCreate{})
}

type ThirdChangeExternalChatCreate struct {
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

func (ThirdChangeExternalChatCreate) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalChatCreate) GetEventType() string {
	return "change_external_chat"
}

func (ThirdChangeExternalChatCreate) GetChangeType() string {
	return "create"
}

func (m ThirdChangeExternalChatCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalChatCreate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalChatCreate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
