package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95804#支付成功通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdLicensePaySuccess{})
}

type ThirdLicensePaySuccess struct {
	XMLName       xml.Name `xml:"xml"`
	Text          string   `xml:",chardata"`
	ServiceCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"ServiceCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	OrderId struct {
		Text string `xml:",chardata"`
	} `xml:"OrderId"`
	BuyerUserId struct {
		Text string `xml:",chardata"`
	} `xml:"BuyerUserId"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
}

func (ThirdLicensePaySuccess) GetMessageType() string {
	return "third"
}

func (ThirdLicensePaySuccess) GetEventType() string {
	return "license_pay_success"
}

func (ThirdLicensePaySuccess) GetChangeType() string {
	return ""
}

func (m ThirdLicensePaySuccess) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdLicensePaySuccess) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdLicensePaySuccess
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
