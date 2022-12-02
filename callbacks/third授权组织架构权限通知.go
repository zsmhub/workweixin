package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/97378#授权组织架构权限通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdCorpArchAuth{})
}

type ThirdCorpArchAuth struct {
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
}

func (ThirdCorpArchAuth) GetMessageType() string {
	return "third"
}

func (ThirdCorpArchAuth) GetEventType() string {
	return "corp_arch_auth"
}

func (ThirdCorpArchAuth) GetChangeType() string {
	return ""
}

func (m ThirdCorpArchAuth) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdCorpArchAuth) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdCorpArchAuth
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
