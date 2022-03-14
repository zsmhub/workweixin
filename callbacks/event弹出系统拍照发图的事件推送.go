package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#弹出系统拍照发图的事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventPicSysphoto{})
}

type EventPicSysphoto struct {
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
	SendPicsInfo struct {
		Text  string `xml:",chardata"`
		Count struct {
			Text string `xml:",chardata"`
		} `xml:"Count"`
		PicList struct {
			Text string `xml:",chardata"`
			Item struct {
				Text      string `xml:",chardata"`
				PicMd5Sum struct {
					Text string `xml:",chardata"`
				} `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventPicSysphoto) GetMessageType() string {
	return "event"
}

func (EventPicSysphoto) GetEventType() string {
	return "pic_sysphoto"
}

func (EventPicSysphoto) GetChangeType() string {
	return ""
}

func (m EventPicSysphoto) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventPicSysphoto) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventPicSysphoto
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
