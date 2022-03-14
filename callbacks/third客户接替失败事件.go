package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#客户接替失败事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactTransferFail{})
}

type ThirdChangeExternalContactTransferFail struct {
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
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
	FailReason struct {
		Text string `xml:",chardata"`
	} `xml:"FailReason"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	ExternalUserID struct {
		Text string `xml:",chardata"`
	} `xml:"ExternalUserID"`
}

func (ThirdChangeExternalContactTransferFail) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactTransferFail) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactTransferFail) GetChangeType() string {
	return "transfer_fail"
}

func (m ThirdChangeExternalContactTransferFail) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactTransferFail) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactTransferFail
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
