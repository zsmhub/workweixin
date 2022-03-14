package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#审批状态通知事件

func init() {
	// 添加可解析的回调事件
	supportCallback(EventOpenApprovalChange{})
}

type EventOpenApprovalChange struct {
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
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	ApprovalInfo struct {
		Text    string `xml:",chardata"`
		ThirdNo struct {
			Text string `xml:",chardata"`
		} `xml:"ThirdNo"`
		OpenSpName struct {
			Text string `xml:",chardata"`
		} `xml:"OpenSpName"`
		OpenTemplateId struct {
			Text string `xml:",chardata"`
		} `xml:"OpenTemplateId"`
		OpenSpStatus struct {
			Text string `xml:",chardata"`
		} `xml:"OpenSpStatus"`
		ApplyTime struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyTime"`
		ApplyUserName struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyUserName"`
		ApplyUserId struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyUserId"`
		ApplyUserParty struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyUserParty"`
		ApplyUserImage struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyUserImage"`
		ApprovalNodes struct {
			Text         string `xml:",chardata"`
			ApprovalNode struct {
				Text       string `xml:",chardata"`
				NodeStatus struct {
					Text string `xml:",chardata"`
				} `xml:"NodeStatus"`
				NodeAttr struct {
					Text string `xml:",chardata"`
				} `xml:"NodeAttr"`
				NodeType struct {
					Text string `xml:",chardata"`
				} `xml:"NodeType"`
				Items struct {
					Text string `xml:",chardata"`
					Item struct {
						Text     string `xml:",chardata"`
						ItemName struct {
							Text string `xml:",chardata"`
						} `xml:"ItemName"`
						ItemUserId struct {
							Text string `xml:",chardata"`
						} `xml:"ItemUserId"`
						ItemImage struct {
							Text string `xml:",chardata"`
						} `xml:"ItemImage"`
						ItemStatus struct {
							Text string `xml:",chardata"`
						} `xml:"ItemStatus"`
						ItemSpeech struct {
							Text string `xml:",chardata"`
						} `xml:"ItemSpeech"`
						ItemOpTime struct {
							Text string `xml:",chardata"`
						} `xml:"ItemOpTime"`
					} `xml:"Item"`
				} `xml:"Items"`
			} `xml:"ApprovalNode"`
		} `xml:"ApprovalNodes"`
		NotifyNodes struct {
			Text       string `xml:",chardata"`
			NotifyNode struct {
				Text     string `xml:",chardata"`
				ItemName struct {
					Text string `xml:",chardata"`
				} `xml:"ItemName"`
				ItemUserId struct {
					Text string `xml:",chardata"`
				} `xml:"ItemUserId"`
				ItemImage struct {
					Text string `xml:",chardata"`
				} `xml:"ItemImage"`
			} `xml:"NotifyNode"`
		} `xml:"NotifyNodes"`
		Approverstep struct {
			Text string `xml:",chardata"`
		} `xml:"approverstep"`
	} `xml:"ApprovalInfo"`
}

func (EventOpenApprovalChange) GetMessageType() string {
	return "event"
}

func (EventOpenApprovalChange) GetEventType() string {
	return "open_approval_change"
}

func (EventOpenApprovalChange) GetChangeType() string {
	return ""
}

func (m EventOpenApprovalChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventOpenApprovalChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventOpenApprovalChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
