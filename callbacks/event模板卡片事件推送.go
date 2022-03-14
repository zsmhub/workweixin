package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#模板卡片事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventTemplateCardEvent{})
}

type EventTemplateCardEvent struct {
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
	EventKey struct {
		Text string `xml:",chardata"`
	} `xml:"EventKey"`
	TaskId struct {
		Text string `xml:",chardata"`
	} `xml:"TaskId"`
	CardType struct {
		Text string `xml:",chardata"`
	} `xml:"CardType"`
	ResponseCode struct {
		Text string `xml:",chardata"`
	} `xml:"ResponseCode"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	SelectedItems struct {
		Text         string `xml:",chardata"`
		SelectedItem []struct {
			Text        string `xml:",chardata"`
			QuestionKey struct {
				Text string `xml:",chardata"`
			} `xml:"QuestionKey"`
			OptionIds struct {
				Text     string `xml:",chardata"`
				OptionId []struct {
					Text string `xml:",chardata"`
				} `xml:"OptionId"`
			} `xml:"OptionIds"`
		} `xml:"SelectedItem"`
	} `xml:"SelectedItems"`
}

func (EventTemplateCardEvent) GetMessageType() string {
	return "event"
}

func (EventTemplateCardEvent) GetEventType() string {
	return "template_card_event"
}

func (EventTemplateCardEvent) GetChangeType() string {
	return ""
}

func (m EventTemplateCardEvent) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventTemplateCardEvent) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventTemplateCardEvent
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
