package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/96219#回调异步任务结果

func init() {
	// 添加可解析的回调事件
	supportCallback(EventUploadMediaJobFinish{})
}

type EventUploadMediaJobFinish struct {
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
	JobId struct {
		Text string `xml:",chardata"`
	} `xml:"JobId"`
}

func (EventUploadMediaJobFinish) GetMessageType() string {
	return "event"
}

func (EventUploadMediaJobFinish) GetEventType() string {
	return "upload_media_job_finish"
}

func (EventUploadMediaJobFinish) GetChangeType() string {
	return ""
}

func (m EventUploadMediaJobFinish) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventUploadMediaJobFinish) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventUploadMediaJobFinish
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
