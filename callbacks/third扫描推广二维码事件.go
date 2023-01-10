package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/98071#扫描推广二维码事件

func init() {
	//添加可解析的回调事件
	supportCallback(ThirdEnterRegisterPackage{})
}

// XML was generated 2023-01-10 11:28:56 by insomnia on insomnia.lan.
type ThirdEnterRegisterPackage struct {
	XMLName       xml.Name `xml:"xml"`
	Text          string   `xml:",chardata"`
	ServiceCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"ServiceCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	AuthUserInfo struct {
		Text   string `xml:",chardata"`
		UserId struct {
			Text string `xml:",chardata"`
		} `xml:"UserId"`
	} `xml:"AuthUserInfo"`
	RegisterCode struct {
		Text string `xml:",chardata"`
	} `xml:"RegisterCode"`
	State struct {
		Text string `xml:",chardata"`
	} `xml:"State"`
	TemplateId struct {
		Text string `xml:",chardata"`
	} `xml:"TemplateId"`
}

func (ThirdEnterRegisterPackage) GetMessageType() string {
	return "third"
}

func (ThirdEnterRegisterPackage) GetEventType() string {
	return "enter_register_package"
}

func (ThirdEnterRegisterPackage) GetChangeType() string {
	return ""
}

func (m ThirdEnterRegisterPackage) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdEnterRegisterPackage) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdEnterRegisterPackage
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
