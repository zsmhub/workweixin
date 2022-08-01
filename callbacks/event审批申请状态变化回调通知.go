package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/92633#2.事件格式

func init() {
	//添加可解析的回调事件
	supportCallback(EventSysApprovalChange{})
}

// XML was generated 2022-07-29 14:39:22 by yanglaosan on yanglaosan.
type EventSysApprovalChange struct {
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
		Text string `xml:",chardata"`
		SpNo struct {
			Text string `xml:",chardata"`
		} `xml:"SpNo"`
		SpName struct {
			Text string `xml:",chardata"`
		} `xml:"SpName"`
		SpStatus struct {
			Text string `xml:",chardata"`
		} `xml:"SpStatus"`
		TemplateId struct {
			Text string `xml:",chardata"`
		} `xml:"TemplateId"`
		ApplyTime struct {
			Text string `xml:",chardata"`
		} `xml:"ApplyTime"`
		Applyer struct {
			Text   string `xml:",chardata"`
			UserId struct {
				Text string `xml:",chardata"`
			} `xml:"UserId"`
			Party struct {
				Text string `xml:",chardata"`
			} `xml:"Party"`
		} `xml:"Applyer"`
		SpRecord []struct {
			Text     string `xml:",chardata"`
			SpStatus struct {
				Text string `xml:",chardata"`
			} `xml:"SpStatus"`
			ApproverAttr struct {
				Text string `xml:",chardata"`
			} `xml:"ApproverAttr"`
			Details []struct {
				Text     string `xml:",chardata"`
				Approver struct {
					Text   string `xml:",chardata"`
					UserId struct {
						Text string `xml:",chardata"`
					} `xml:"UserId"`
				} `xml:"Approver"`
				Speech struct {
					Text string `xml:",chardata"`
				} `xml:"Speech"`
				SpStatus struct {
					Text string `xml:",chardata"`
				} `xml:"SpStatus"`
				SpTime struct {
					Text string `xml:",chardata"`
				} `xml:"SpTime"`
			} `xml:"Details"`
		} `xml:"SpRecord"`
		Notifyer struct {
			Text   string `xml:",chardata"`
			UserId struct {
				Text string `xml:",chardata"`
			} `xml:"UserId"`
		} `xml:"Notifyer"`
		Comments struct {
			Text            string `xml:",chardata"`
			CommentUserInfo struct {
				Text   string `xml:",chardata"`
				UserId struct {
					Text string `xml:",chardata"`
				} `xml:"UserId"`
			} `xml:"CommentUserInfo"`
			CommentTime struct {
				Text string `xml:",chardata"`
			} `xml:"CommentTime"`
			CommentContent struct {
				Text string `xml:",chardata"`
			} `xml:"CommentContent"`
			CommentId struct {
				Text string `xml:",chardata"`
			} `xml:"CommentId"`
		} `xml:"Comments"`
		StatuChangeEvent struct {
			Text string `xml:",chardata"`
		} `xml:"StatuChangeEvent"`
	} `xml:"ApprovalInfo"`
}

func (EventSysApprovalChange) GetMessageType() string {
	return "event"
}

func (EventSysApprovalChange) GetEventType() string {
	return "sys_approval_change"
}

func (EventSysApprovalChange) GetChangeType() string {
	return ""
}

func (m EventSysApprovalChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventSysApprovalChange) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventSysApprovalChange
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
