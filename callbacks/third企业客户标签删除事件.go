package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#企业客户标签删除事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalTagDelete{})
}

type ThirdChangeExternalTagDelete struct {
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
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
	TagType struct {
		Text string `xml:",chardata"`
	} `xml:"TagType"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (ThirdChangeExternalTagDelete) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalTagDelete) GetEventType() string {
	return "change_external_tag"
}

func (ThirdChangeExternalTagDelete) GetChangeType() string {
	return "delete"
}

func (m ThirdChangeExternalTagDelete) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalTagDelete) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalTagDelete
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
