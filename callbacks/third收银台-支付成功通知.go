package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/91931#支付成功通知

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdPayForAppSuccess{})
}

type ThirdPayForAppSuccess struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	PaidCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"PaidCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	OrderId struct {
		Text string `xml:",chardata"`
	} `xml:"OrderId"`
}

func (ThirdPayForAppSuccess) GetMessageType() string {
	return "third"
}

func (ThirdPayForAppSuccess) GetEventType() string {
	return "pay_for_app_success"
}

func (ThirdPayForAppSuccess) GetChangeType() string {
	return ""
}

func (m ThirdPayForAppSuccess) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdPayForAppSuccess) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdPayForAppSuccess
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
