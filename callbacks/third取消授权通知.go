package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/90642#取消授权通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdCancelAuth{})
}

type ThirdCancelAuth struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
}

func (ThirdCancelAuth) GetMessageType() string {
	return "third"
}

func (ThirdCancelAuth) GetEventType() string {
	return "cancel_auth"
}

func (ThirdCancelAuth) GetChangeType() string {
	return ""
}

func (m ThirdCancelAuth) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdCancelAuth) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdCancelAuth
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
