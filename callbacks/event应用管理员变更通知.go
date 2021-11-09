package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://work.weixin.qq.com/api/doc/90001/90143/95038#应用管理员变更通知

func init() {
    // 添加可解析的回调事件
    supportCallback(EventChangeAppAdmin{})
}

// XML was generated 2021-10-09 14:52:26 by insomnia on Insomnia.lan.
type EventChangeAppAdmin struct {
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

func (EventChangeAppAdmin) GetMessageType() string {
    return "event"
}

func (EventChangeAppAdmin) GetEventType() string {
    return "change_app_admin"
}

func (EventChangeAppAdmin) GetChangeType() string {
    return ""
}

func (m EventChangeAppAdmin) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventChangeAppAdmin) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp EventChangeAppAdmin
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
