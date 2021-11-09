package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90375#文本消息

func init() {
    // 添加可解析的回调事件
    supportCallback(Text{})
}

// XML was generated 2021-11-01 20:28:46 by insomnia on Insomnia.lan.
type Text struct {
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
    Content struct {
        Text string `xml:",chardata"`
    } `xml:"Content"`
    MsgId struct {
        Text string `xml:",chardata"`
    } `xml:"MsgId"`
    AgentID struct {
        Text string `xml:",chardata"`
    } `xml:"AgentID"`
}

func (Text) GetMessageType() string {
    return "text"
}

func (Text) GetEventType() string {
    return ""
}

func (Text) GetChangeType() string {
    return ""
}

func (m Text) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (Text) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp Text
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
