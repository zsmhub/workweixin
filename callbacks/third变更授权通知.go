package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/90642#变更授权通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeAuth{})
}

type ThirdChangeAuth struct {
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

func (ThirdChangeAuth) GetMessageType() string {
	return "third"
}

func (ThirdChangeAuth) GetEventType() string {
	return "change_auth"
}

func (ThirdChangeAuth) GetChangeType() string {
	return ""
}

func (m ThirdChangeAuth) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeAuth) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeAuth
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
