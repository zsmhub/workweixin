package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#客户群变更事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalChatUpdate{})
}

type ThirdChangeExternalChatUpdate struct {
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
	UpdateDetail struct {
		Text string `xml:",chardata"`
	} `xml:"UpdateDetail"`
	JoinScene struct {
		Text string `xml:",chardata"`
	} `xml:"JoinScene"`
	QuitScene struct {
		Text string `xml:",chardata"`
	} `xml:"QuitScene"`
	MemChangeCnt struct {
		Text string `xml:",chardata"`
	} `xml:"MemChangeCnt"`
}

func (ThirdChangeExternalChatUpdate) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalChatUpdate) GetEventType() string {
	return "change_external_chat"
}

func (ThirdChangeExternalChatUpdate) GetChangeType() string {
	return "update"
}

func (m ThirdChangeExternalChatUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalChatUpdate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalChatUpdate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
