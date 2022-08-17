package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/94758#重置永久授权码通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdResetPermanentCode{})
}

type ThirdResetPermanentCode struct {
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
}

func (ThirdResetPermanentCode) GetMessageType() string {
	return "third"
}

func (ThirdResetPermanentCode) GetEventType() string {
	return "reset_permanent_code"
}

func (ThirdResetPermanentCode) GetChangeType() string {
	return ""
}

func (m ThirdResetPermanentCode) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdResetPermanentCode) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdResetPermanentCode
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
