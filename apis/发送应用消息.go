package apis

import (
	"bytes"
	"encoding/json"
)

// 发送应用消息
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义

// 文本消息
type ReqSentTextBody struct {
	Content string `json:"content"` // 消息内容，最长不超过2048个字节，超过将截断（支持id转译）
}

// 图片消息
type ReqSentImageBody struct {
	MediaId string `json:"media_id"` // 图片媒体文件id，可以调用上传临时素材接口获取
}

// 语音消息
type ReqSentVoiceBody struct {
	MediaId string `json:"media_id"` // 语音文件id，可以调用上传临时素材接口获取
}

// 视频消息
type ReqSentVideoBody struct {
	MediaId     string `json:"media_id"`    // 视频媒体文件id，可以调用上传临时素材接口获取
	Title       string `json:"title"`       // 视频消息的标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 视频消息的描述，不超过512个字节，超过会自动截断
}

// 文件消息
type ReqSentFileBody struct {
	MediaId string `json:"media_id"` // 文件id，可以调用上传临时素材接口获取
}

// 文本卡片
type ReqSentMessageCardTextBody struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。 最长2048字节，请确保包含了协议头(http/https)，小程序或者url必须填写一个
	BtnTxt      string `json:"btntxt"`      // 按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断。
}

// 图文消息
type ReqSentMessageCardNewsArticleBody struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断（支持id转译）
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断（支持id转译）
	Url         string `json:"url"`         // 点击后跳转的链接。 最长2048字节，请确保包含了协议头(http/https)，小程序或者url必须填写一个
	UrlImg      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
	Appid       string `json:"appid"`       // 小程序appid，必须是与当前应用关联的小程序，appid和pagepath必须同时填写，填写后会忽略url字段
	Pagepath    string `json:"pagepath"`    // 点击消息卡片后的小程序页面，仅限本小程序内的页面。appid和pagepath必须同时填写，填写后会忽略url字段
}
type ReqSentMessageCardNewsBody struct {
	Articles []ReqSentMessageCardNewsArticleBody `json:"articles"`
}

// 图文消息（mpnews）
type ReqSentMpNewsArticle struct {
	Title            string `json:"title"`              // 标题，不超过128个字节，超过会自动截断（支持id转译）
	ThumbMediaId     string `json:"thumb_media_id"`     // 图文消息缩略图的media_id, 可以通过素材管理接口获得。此处thumb_media_id即上传接口返回的media_id
	Author           string `json:"author"`             // 图文消息的作者，不超过64个字节
	ContentSourceUrl string `json:"content_source_url"` // 图文消息点击“阅读原文”之后的页面链接
	Content          string `json:"content"`            // 图文消息的内容，支持html标签，不超过666 K个字节（支持id转译）
	Digest           string `json:"digest"`             // 图文消息的描述，不超过512个字节，超过会自动截断（支持id转译）
}
type ReqSentMpNewsBody struct {
	Articles []ReqSentMpNewsArticle `json:"articles"`
}

// markdown消息
type ReqSentMarkdownBody struct {
	Content string `json:"content"` // markdown内容，最长不超过2048个字节，必须是utf8编码
}

type ReqSentMessageCard struct {
	ToUser                 string                     `json:"touser,omitempty"`                   // 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToParty                string                     `json:"toparty,omitempty"`                  // 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag                  string                     `json:"totag,omitempty"`                    // 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	MsgType                string                     `json:"msgtype"`                            // 消息类型，如text/image/file
	AgentId                int                        `json:"agentid"`                            // 企业应用的id，整型。企业内部开发，可在应用的设置页面查看；第三方服务商，可通过接口 获取企业授权信息 获取该参数值
	EnableIdTrans          int                        `json:"enable_id_trans,omitempty"`          // 表示是否开启id转译，0表示否，1表示是，默认0
	EnableDuplicateCheck   int                        `json:"enable_duplicate_check,omitempty"`   // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	DuplicateCheckInterval int                        `json:"duplicate_check_interval,omitempty"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	Safe                   int                        `json:"safe,omitempty"`                     // 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，2表示仅限在企业内分享，默认为0；注意仅mpnews类型的消息支持safe值为2，其他消息类型不支持
	Text                   ReqSentTextBody            `json:"text"`
	Image                  ReqSentImageBody           `json:"image"`
	Voice                  ReqSentVoiceBody           `json:"voice"`
	Video                  ReqSentVideoBody           `json:"video"`
	File                   ReqSentFileBody            `json:"file"`
	Textcard               ReqSentMessageCardTextBody `json:"textcard"`
	News                   ReqSentMessageCardNewsBody `json:"news"`
	Mpnews                 ReqSentMpNewsArticle       `json:"mpnews"`
	Markdown               ReqSentMarkdownBody        `json:"markdown"`
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

// RespSentMessageCard 接口定义响应
// 文档：https://developer.work.weixin.qq.com/document/path/90372#接口定义
type RespSentMessageCard struct {
	CommonResp
	InvalidUser  string `json:"invaliduser"`   // 不合法的userid
	InvalidParty string `json:"invalidparty"`  // 不合法的partyid
	InvalidTag   string `json:"invalidtag"`    // 不合法的标签id
	MsgId        string `json:"msgid"`         // 消息id，用于撤回应用消息
	ResponseCode string `json:"response_code"` // 仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回，应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
}

// 发送应用消息
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
