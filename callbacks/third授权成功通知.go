package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/90642#授权成功通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdCreateAuth{})
}

type ThirdCreateAuth struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	AuthCode struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCode"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	State struct {
		Text string `xml:",chardata"`
	} `xml:"State"`
}

func (ThirdCreateAuth) GetMessageType() string {
	return "third"
}

func (ThirdCreateAuth) GetEventType() string {
	return "create_auth"
}

func (ThirdCreateAuth) GetChangeType() string {
	return ""
}

func (m ThirdCreateAuth) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdCreateAuth) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdCreateAuth
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
