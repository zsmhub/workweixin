package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://open.work.weixin.qq.com/api/doc/90001/90143/90628#推送suite_ticket

func init() {
    // 添加可解析的回调事件
    supportCallback(ThirdSuiteTicket{})
}

// XML was generated 2021-09-16 10:26:42 by chenjianlin on JZTech-chenjianlin.lan.
type ThirdSuiteTicket struct {
    XMLName xml.Name `xml:"xml"`
    Text    string   `xml:",chardata"`
    SuiteId struct {
        Text string `xml:",chardata"`
    } `xml:"SuiteId"`
    InfoType struct {
        Text string `xml:",chardata"`
    } `xml:"InfoType"`
    TimeStamp struct {
        Text string `xml:",chardata"`
    } `xml:"TimeStamp"`
    SuiteTicket struct {
        Text string `xml:",chardata"`
    } `xml:"SuiteTicket"`
}

func (ThirdSuiteTicket) GetMessageType() string {
    return "third"
}

func (ThirdSuiteTicket) GetEventType() string {
    return "suite_ticket"
}

func (ThirdSuiteTicket) GetChangeType() string {
    return ""
}

func (m ThirdSuiteTicket) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdSuiteTicket) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp ThirdSuiteTicket
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
