package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90970#新增成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventChangeContactCreateUser{})
}

type EventChangeContactCreateUser struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	FromUserName struct {
		Text string `xml:",chardata"`
	} `xml:"FromUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"Name"`
	Department struct {
		Text string `xml:",chardata"`
	} `xml:"Department"`
	MainDepartment struct {
		Text string `xml:",chardata"`
	} `xml:"MainDepartment"`
	IsLeaderInDept struct {
		Text string `xml:",chardata"`
	} `xml:"IsLeaderInDept"`
	DirectLeader struct {
		Text string `xml:",chardata"`
	} `xml:"DirectLeader"`
	Position struct {
		Text string `xml:",chardata"`
	} `xml:"Position"`
	Mobile struct {
		Text string `xml:",chardata"`
	} `xml:"Mobile"`
	Gender struct {
		Text string `xml:",chardata"`
	} `xml:"Gender"`
	Email struct {
		Text string `xml:",chardata"`
	} `xml:"Email"`
	BizMail struct {
		Text string `xml:",chardata"`
	} `xml:"BizMail"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"Status"`
	Avatar struct {
		Text string `xml:",chardata"`
	} `xml:"Avatar"`
	Alias struct {
		Text string `xml:",chardata"`
	} `xml:"Alias"`
	Telephone struct {
		Text string `xml:",chardata"`
	} `xml:"Telephone"`
	Address struct {
		Text string `xml:",chardata"`
	} `xml:"Address"`
	ExtAttr struct {
		Text string `xml:",chardata"`
		Item []struct {
			Chardata string `xml:",chardata"`
			Name     struct {
				Text string `xml:",chardata"`
			} `xml:"Name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"Type"`
			Text struct {
				Text  string `xml:",chardata"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"Value"`
			} `xml:"Text"`
			Web struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"`
				} `xml:"Title"`
				URL struct {
					Text string `xml:",chardata"`
				} `xml:"Url"`
			} `xml:"Web"`
		} `xml:"Item"`
	} `xml:"ExtAttr"`
}

func (EventChangeContactCreateUser) GetMessageType() string {
	return "event"
}

func (EventChangeContactCreateUser) GetEventType() string {
	return "change_contact"
}

func (EventChangeContactCreateUser) GetChangeType() string {
	return "create_user"
}

func (m EventChangeContactCreateUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeContactCreateUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventChangeContactCreateUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
