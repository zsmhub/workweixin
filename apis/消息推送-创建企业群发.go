package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqAddMsgTemplateExternalcontact 创建企业群发请求
// 文档：https://developer.work.weixin.qq.com/document/path/92698#创建企业群发
type (
	AddMsgTemplateText struct {
		// Content 消息文本内容，最多4000个<strong>字节</strong>
		Content string `json:"content"`
	}

	AddMsgTemplateAttachments struct {
		File        AddMsgTemplateAttachmentsFile        `json:"file"`
		Image       AddMsgTemplateAttachmentsImage       `json:"image"`
		Link        AddMsgTemplateAttachmentsLink        `json:"link"`
		Miniprogram AddMsgTemplateAttachmentsMiniprogram `json:"miniprogram"`
		// Msgtype 附件类型，可选image、link、miniprogram或者video，必填
		Msgtype string                         `json:"msgtype"`
		Video   AddMsgTemplateAttachmentsVideo `json:"video"`
	}

	AddMsgTemplateAttachmentsFile struct {
		MediaID string `json:"media_id"`
	}

	AddMsgTemplateAttachmentsImage struct {
		MediaID string `json:"media_id"`
		// PicURL 图片的链接，<strong>仅可使用<a href="#13219">上传图片</a>接口得到的链接</strong>
		PicURL string `json:"pic_url"`
	}

	AddMsgTemplateAttachmentsLink struct {
		// Desc 图文消息的描述，最多512个<strong>字节</strong>
		Desc string `json:"desc"`
		// Picurl 图文消息封面的url，最长2048个字节
		Picurl string `json:"picurl"`
		// Title 图文消息标题，最长128个字节，必填
		// Title 小程序消息标题，最多64个<strong>字节</strong>，必填
		Title string `json:"title"`
		// URL 图文消息的链接，最长2048个字节，必填
		URL string `json:"url"`
	}

	AddMsgTemplateAttachmentsMiniprogram struct {
		// Appid 小程序appid（可以在微信公众平台上查询），必须是<strong>关联到企业的小程序应用</strong>，必填
		Appid string `json:"appid"`
		// Page 小程序page路径，必填
		Page string `json:"page"`
		// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416，必填
		PicMediaID string `json:"pic_media_id"`
		// Title 图文消息标题，最长128个字节，必填
		// Title 小程序消息标题，最多64个<strong>字节</strong>，必填
		Title string `json:"title"`
	}

	AddMsgTemplateAttachmentsVideo struct {
		MediaID string `json:"media_id"`
	}
)

type ReqAddMsgTemplateExternalcontact struct {
	Attachments []AddMsgTemplateAttachments `json:"attachments"` // 附件，最多支持添加9个附件
	// ChatType 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	ChatType string `json:"chat_type"`
	// ExternalUserid 客户的外部联系人id列表，仅在chat_type为single时有效，不可与sender同时为空，最多可传入<strong>1万</strong>个客户
	ExternalUserid []string `json:"external_userid"`
	// Sender 发送企业群发消息的成员userid，当类型为发送给客户群时必填
	Sender string             `json:"sender"`
	Text   AddMsgTemplateText `json:"text"`
	// 是否允许成员在待发送客户列表中重新进行选择，默认为false
	AllowSelect bool `json:"allow_select"`
}

var _ bodyer = ReqAddMsgTemplateExternalcontact{}

func (x ReqAddMsgTemplateExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespAddMsgTemplateExternalcontact 创建企业群发响应
// 文档：https://developer.work.weixin.qq.com/document/path/92698#创建企业群发
type RespAddMsgTemplateExternalcontact struct {
	CommonResp
	FailList []string `json:"fail_list"`
	Msgid    string   `json:"msgid"`
}

var _ bodyer = RespAddMsgTemplateExternalcontact{}

func (x RespAddMsgTemplateExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddMsgTemplateExternalcontact 创建企业群发
// 文档：https://developer.work.weixin.qq.com/document/path/92698#创建企业群发
func (c *ApiClient) ExecAddMsgTemplateExternalcontact(req ReqAddMsgTemplateExternalcontact) (RespAddMsgTemplateExternalcontact, error) {
	var resp RespAddMsgTemplateExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/add_msg_template", req, &resp, true)
	if err != nil {
		return RespAddMsgTemplateExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddMsgTemplateExternalcontact{}, bizErr
	}
	return resp, nil
}
