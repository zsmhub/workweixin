package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/95994#自动激活回调通知

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdAutoActivate{})
}

type ThirdAutoActivate struct {
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
	Scene struct {
		Text string `xml:",chardata"`
	} `xml:"Scene"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	AccountList struct {
		Text       string `xml:",chardata"`
		ActiveCode struct {
			Text string `xml:",chardata"`
		} `xml:"ActiveCode"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"Type"`
		ExpireTime struct {
			Text string `xml:",chardata"`
		} `xml:"ExpireTime"`
		UserId struct {
			Text string `xml:",chardata"`
		} `xml:"UserId"`
		PreviousStatus struct {
			Text string `xml:",chardata"`
		} `xml:"PreviousStatus"`
		PreviousActiveCode struct {
			Text string `xml:",chardata"`
		} `xml:"PreviousActiveCode"`
	} `xml:"AccountList"`
}

func (ThirdAutoActivate) GetMessageType() string {
	return "third"
}

func (ThirdAutoActivate) GetEventType() string {
	return "auto_activate"
}

func (ThirdAutoActivate) GetChangeType() string {
	return ""
}

func (m ThirdAutoActivate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdAutoActivate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdAutoActivate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
