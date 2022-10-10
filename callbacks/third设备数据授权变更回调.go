package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/96103#设备数据授权变更回调

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdDeviceDataAuthChange{})
}

type ThirdDeviceDataAuthChange struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (ThirdDeviceDataAuthChange) GetMessageType() string {
	return "third"
}

func (ThirdDeviceDataAuthChange) GetEventType() string {
	return "device_data_auth_change"
}

func (ThirdDeviceDataAuthChange) GetChangeType() string {
	return ""
}

func (m ThirdDeviceDataAuthChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdDeviceDataAuthChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdDeviceDataAuthChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
