package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90585#注册完成回调事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdRegisterCorp{})
}

type ThirdRegisterCorp struct {
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
	RegisterCode struct {
		Text string `xml:",chardata"`
	} `xml:"RegisterCode"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	ContactSync struct {
		Text        string `xml:",chardata"`
		AccessToken struct {
			Text string `xml:",chardata"`
		} `xml:"AccessToken"`
		ExpiresIn struct {
			Text string `xml:",chardata"`
		} `xml:"ExpiresIn"`
	} `xml:"ContactSync"`
	AuthUserInfo struct {
		Text   string `xml:",chardata"`
		UserId struct {
			Text string `xml:",chardata"`
		} `xml:"UserId"`
	} `xml:"AuthUserInfo"`
	State struct {
		Text string `xml:",chardata"`
	} `xml:"State"`
	TemplateId struct {
		Text string `xml:",chardata"`
	} `xml:"TemplateId"`
}

func (ThirdRegisterCorp) GetMessageType() string {
	return "third"
}

func (ThirdRegisterCorp) GetEventType() string {
	return "register_corp"
}

func (ThirdRegisterCorp) GetChangeType() string {
	return ""
}

func (m ThirdRegisterCorp) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdRegisterCorp) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdRegisterCorp
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
