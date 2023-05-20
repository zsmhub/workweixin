package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/98959#获客助手权限取消事件

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdCancelSpecialAuth{})
}

type ThirdCancelSpecialAuth struct {
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
	AuthType struct {
		Text string `xml:",chardata"`
	} `xml:"AuthType"`
}

func (ThirdCancelSpecialAuth) GetMessageType() string {
	return "third"
}

func (ThirdCancelSpecialAuth) GetEventType() string {
	return "cancel_special_auth"
}

func (ThirdCancelSpecialAuth) GetChangeType() string {
	return ""
}

func (m ThirdCancelSpecialAuth) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdCancelSpecialAuth) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdCancelSpecialAuth
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
