package apis

import (
	"bytes"
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSentMessage 接口定义请求
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

// 文本卡片
type ReqSentMessageCardTextBody struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。 最长2048字节，请确保包含了协议头(http/https)，小程序或者url必须填写一个
	BtnTxt      string `json:"btntxt"`      // 按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断。
}

// 图文消息
type ReqSentMessageCardNewsArticleBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlImg      string `json:"picurl"`
}
type ReqSentMessageCardNewsBody struct {
	Articles []ReqSentMessageCardNewsArticleBody `json:"articles"`
}

type ReqSentMessageCard struct {
	ToUser                 string                     `json:"touser"`  // 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToParty                string                     `json:"toparty"` // 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag                  string                     `json:"totag"`   // 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	MsgType                string                     `json:"msgtype"`
	AgentId                int                        `json:"agentid"`                  // 企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
	EnableIdTrans          string                     `json:"enable_id_trans"`          // 表示是否开启id转译，0表示否，1表示是，默认0
	EnableDuplicateCheck   string                     `json:"enable_duplicate_check"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval string                     `json:"duplicate_check_interval"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	Textcard               ReqSentMessageCardTextBody `json:"textcard"`
	News                   ReqSentMessageCardNewsBody `json:"news"`
}

var _ bodyer = ReqSentMessageCard{}

func (x ReqSentMessageCard) intoBody() ([]byte, error) {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(x)
	if err != nil {
		panic(err)
	}
	return byteBuf.Bytes(), nil
}

// RespSentMessage 接口定义响应
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

type RespSentMessageCard struct {
	CommonResp
	InvalidUser  string `json:"invaliduser"`   // 不合法的userid
	InvalidParty string `json:"invalidparty"`  // 不合法的partyid
	InvalidTag   string `json:"invalidtag"`    // 不合法的标签id
	MsgId        string `json:"msgid"`         // 消息id，用于撤回应用消息
	ResponseCode string `json:"response_code"` // 仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回，应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
}

// execSentMessage 接口定义
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义
func (c *ApiClient) ExecSentMessageCard(req ReqSentMessageCard) (RespSentMessageCard, error) {
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
