package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92654#新增成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeContactCreateUser{})
}

type ThirdChangeContactCreateUser struct {
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
	OpenUserID struct {
		Text string `xml:",chardata"`
	} `xml:"OpenUserID"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"Name"`
	Department struct {
		Text string `xml:",chardata"`
	} `xml:"Department"`
	Gender struct {
		Text string `xml:",chardata"`
	} `xml:"Gender"`
	Avatar struct {
		Text string `xml:",chardata"`
	} `xml:"Avatar"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"Status"`
	IsLeaderInDept struct {
		Text string `xml:",chardata"`
	} `xml:"IsLeaderInDept"`
	MainDepartment struct {
		Text string `xml:",chardata"`
	} `xml:"MainDepartment"`
}

func (ThirdChangeContactCreateUser) GetMessageType() string {
	return "third"
}

func (ThirdChangeContactCreateUser) GetEventType() string {
	return "change_contact"
}

func (ThirdChangeContactCreateUser) GetChangeType() string {
	return "create_user"
}

func (m ThirdChangeContactCreateUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeContactCreateUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeContactCreateUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
