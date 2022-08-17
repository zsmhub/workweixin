package apis

import (
	"bytes"
	"encoding/json"
)

// 小程序通知消息：发送小程序通知消息需要是关联的小程序应用获取的token才能调用，普通的应用不支持发送小程序通知消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

type (
	SentMiniNoticeContentItem struct {
		Key   string `json:"key"`   // 长度10个汉字以内
		Value string `json:"value"` // 长度30个汉字以内（支持id转译）
	}

	SentMiniNoticeBody struct {
		Appid             string                      `json:"appid"`               // 小程序appid，必须是与当前应用关联的小程序
		Page              string                      `json:"page"`                // 点击消息卡片后的小程序页面，仅限本小程序内的页面。该字段不填则消息点击后不跳转。
		Title             string                      `json:"title"`               // 消息标题，长度限制4-12个汉字（支持id转译）
		Description       string                      `json:"description"`         // 消息描述，长度限制4-12个汉字（支持id转译）
		EmphasisFirstItem bool                        `json:"emphasis_first_item"` // 是否放大第一个content_item
		ContentItem       []SentMiniNoticeContentItem `json:"content_item"`        // 消息内容键值对，最多允许10个item
	}
)

type ReqSentMessageCardMiniNotice struct {
	ToUser                 string             `json:"touser,omitempty"`                   // 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToParty                string             `json:"toparty,omitempty"`                  // 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag                  string             `json:"totag,omitempty"`                    // 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	MsgType                string             `json:"msgtype"`                            // 消息类型，如text/image/file
	EnableIdTrans          string             `json:"enable_id_trans,omitempty"`          // 表示是否开启id转译，0表示否，1表示是，默认0
	EnableDuplicateCheck   string             `json:"enable_duplicate_check,omitempty"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval string             `json:"duplicate_check_interval,omitempty"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	MiniprogramNotice      SentMiniNoticeBody `json:"miniprogram_notice"`
}

var _ bodyer = ReqSentMessageCardMiniNotice{}

func (x ReqSentMessageCardMiniNotice) intoBody() ([]byte, error) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(x)
	if err != nil {
		panic(err)
	}
	return byteBuf.Bytes(), nil
}

// 发送应用消息-小程序通知消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义
func (c *ApiClient) ExecSentMessageCardMiniNotice(req ReqSentMessageCardMiniNotice) (RespSentMessageCard, error) {
	var resp RespSentMessageCard
	err := c.executeWXApiPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return RespSentMessageCard{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSentMessageCard{}, bizErr
	}
	return resp, nil
}
