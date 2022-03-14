package callbacks

import "encoding/xml"

// 文档地址
// https://open.work.weixin.qq.com/api/doc/35747#%E5%90%8C%E6%84%8F%E6%8E%88%E6%9D%83%E8%BD%AC%E6%8D%A2external_userid%E4%BA%8B%E4%BB%B6

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdAgreeExternalUseridMigration{})
}

type ThirdAgreeExternalUseridMigration struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	ServiceCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"ServiceCorpId"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
}

func (ThirdAgreeExternalUseridMigration) GetMessageType() string {
	return "third"
}

func (ThirdAgreeExternalUseridMigration) GetEventType() string {
	return "agree_external_userid_migration"
}

func (ThirdAgreeExternalUseridMigration) GetChangeType() string {
	return ""
}

func (m ThirdAgreeExternalUseridMigration) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdAgreeExternalUseridMigration) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdAgreeExternalUseridMigration
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
