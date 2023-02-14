package callbacks

import (
	"encoding/xml"
	"errors"
	"strings"
)

// CallbackMessage 一条接收到的消息
type CallbackMessage struct {
	// ToUserName 企业微信CorpID
	ToUserName string `xml:"ToUserName"` // 企业id
	// FromUserName 成员UserID
	FromUserName string `xml:"FromUserName"` // 固定为sys, 表示系统生成的消息
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// MsgType 消息类型
	MsgType MessageType `xml:"MsgType"`
	// MsgID 消息id，64位整型
	MsgID int64 `xml:"MsgId"`
	// AppAgentId 企业应用的id，整型。可在应用的设置页面查看
	AgentID int64 `xml:"AppAgentId"`
	// EventType 事件类型 MsgType为event存在
	EventType EventType `xml:"Event"`
	// ChangeType 变更类型 Event为change_external_contact存在
	ChangeType ChangeType `xml:"ChangeType"` // 自建应用和第三方应用共用字段

	SuiteId    string    `xml:"SuiteId"`    // 第三方应用ID, 等同于FromUserName
	AuthCorpId string    `xml:"AuthCorpId"` // 授权企业的CorpID, 等同于ToUserName
	InfoType   EventType `xml:"InfoType"`   // 等同于EventType
	TimeStamp  int64     `xml:"TimeStamp"`  // 时间戳

	// 额外的信息
	Extras CallBackExtraInfoInterface

	// 保留原始回调数据，方便排查问题
	OriginalMessage string
}

func (m CallbackMessage) ParseMessageFromXml(body []byte) (CallbackMessage, error) {
	err := xml.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}

	m.OriginalMessage = string(body)

	if m.MsgType == "" {
		m.MsgType = MessageTypeThird
	}
	if m.EventType == "" {
		m.EventType = m.InfoType
	}
	m.EventType = EventType(strings.Trim(string(m.EventType), " "))

	if m.CreateTime == 0 {
		m.CreateTime = m.TimeStamp
	}
	if m.ToUserName == "" {
		m.ToUserName = m.AuthCorpId
	}

	extraParser, ok := callBackParseExtraInfoMap[m.GetTypeKey()]
	if ok {
		m.Extras, err = extraParser.ParseFromXml(body)
	} else {
		err = errors.New("回调事件解析失败，去SDK生成对应的回调事件：" + m.GetTypeKey() + "，xml: " + m.OriginalMessage)
	}
	return m, err
}

func (m CallbackMessage) GetTypeKey() string {
	return string(m.MsgType) + ":" + string(m.EventType) + ":" + string(m.ChangeType)
}

func (m CallbackMessage) GetStructName() string {
	return m.snakeToCamel(string(m.MsgType)) + m.snakeToCamel(string(m.EventType)) + m.snakeToCamel(string(m.ChangeType))
}

func (m CallbackMessage) snakeToCamel(s string) string {
	news := ""
	for k, v := range s {
		if k == 0 {
			news += strings.ToUpper(string(v))
			continue
		}

		if v == '_' {
			continue
		}

		if s[k-1] == '_' {
			news += strings.ToUpper(string(v))
		} else {
			news += string(v)
		}

	}
	return news
}
