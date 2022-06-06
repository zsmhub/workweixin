package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95805#退款结果通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdLicenseRefund{})
}

type ThirdLicenseRefund struct {
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
	OrderStatus struct {
		Text string `xml:",chardata"`
	} `xml:"OrderStatus"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
}

func (ThirdLicenseRefund) GetMessageType() string {
	return "third"
}

func (ThirdLicenseRefund) GetEventType() string {
	return "license_refund"
}

func (ThirdLicenseRefund) GetChangeType() string {
	return ""
}

func (m ThirdLicenseRefund) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdLicenseRefund) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdLicenseRefund
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
