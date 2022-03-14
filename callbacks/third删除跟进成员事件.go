package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#删除跟进成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalContactDelFollowUser{})
}

type ThirdChangeExternalContactDelFollowUser struct {
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
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	ExternalUserID struct {
		Text string `xml:",chardata"`
	} `xml:"ExternalUserID"`
}

func (ThirdChangeExternalContactDelFollowUser) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalContactDelFollowUser) GetEventType() string {
	return "change_external_contact"
}

func (ThirdChangeExternalContactDelFollowUser) GetChangeType() string {
	return "del_follow_user"
}

func (m ThirdChangeExternalContactDelFollowUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalContactDelFollowUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalContactDelFollowUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
