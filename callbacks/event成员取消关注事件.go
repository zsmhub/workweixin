package callbacks

import "encoding/xml"

// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90376#成员关注及取消关注事件

func init() {
    // 添加可解析的回调事件
    supportCallback(EventUnsubscribe{})
}

// XML was generated 2021-10-09 14:46:10 by insomnia on Insomnia.lan.
type EventUnsubscribe struct {
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
}

func (EventUnsubscribe) GetMessageType() string {
    return "event"
}

func (EventUnsubscribe) GetEventType() string {
    return "unsubscribe"
}

func (EventUnsubscribe) GetChangeType() string {
    return ""
}

func (m EventUnsubscribe) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventUnsubscribe) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp EventUnsubscribe
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
