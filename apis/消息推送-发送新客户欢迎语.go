package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqSendWelcomeMsgExternalcontact 发送新客户欢迎语请求
// 文档：https://developer.work.weixin.qq.com/document/path/92599#发送新客户欢迎语
type (
	ReqSendWelcomeMsgExternalcontactFile struct {
		// MediaID 图片的media_id，可以通过素材管理接口获得
		MediaID string `json:"media_id,omitempty"`
	}

	ReqSendWelcomeMsgExternalcontactImage struct {
		// MediaID 视频的media_id，可以通过素材管理接口获得，必填
		// MediaID 文件的media_id,  可以通过素材管理接口获得，必填
		// MediaID 图片的media_id，可以通过素材管理接口获得
		MediaID string `json:"media_id,omitempty"`
		// PicURL 图片的链接，仅可使用上传图片接口得到的链接
		PicURL string `json:"pic_url,omitempty"`
	}
	ReqSendWelcomeMsgExternalcontactLink struct {
		// Desc 图文消息的描述，最长为512字节
		Desc string `json:"desc,omitempty"`
		// Picurl 图文消息封面的url
		Picurl string `json:"picurl,omitempty"`
		// Title 图文消息标题，最长为128字节，必填
		// Title 小程序消息标题，最长为64字节，必填
		Title string `json:"title"`
		// URL 图文消息的链接，必填
		URL string `json:"url"`
	}

	ReqSendWelcomeMsgExternalcontactMiniprogram struct {
		// Appid 小程序appid，必须是关联到企业的小程序应用，必填
		Appid string `json:"appid"`
		// Page 小程序page路径，必填
		Page string `json:"page"`
		// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416，必填
		PicMediaID string `json:"pic_media_id"`
		// Title 图文消息标题，最长为128字节，必填
		// Title 小程序消息标题，最长为64字节，必填
		Title string `json:"title"`
	}

	ReqSendWelcomeMsgExternalcontactVideo struct {
		// MediaID 视频的media_id，可以通过素材管理接口获得，必填
		// MediaID 文件的media_id,  可以通过素材管理接口获得，必填
		// MediaID 图片的media_id，可以通过素材管理接口获得
		MediaID string `json:"media_id,omitempty"`
	}

	ReqSendWelcomeMsgExternalcontactText struct {
		// Content 消息文本内容,最长为4000字节
		Content string `json:"content,omitempty"`
	}
)

type ReqSendWelcomeMsgExternalcontact struct {
	Attachments []struct {
		File        ReqSendWelcomeMsgExternalcontactFile        `json:"file"`
		Image       ReqSendWelcomeMsgExternalcontactImage       `json:"image"`
		Link        ReqSendWelcomeMsgExternalcontactLink        `json:"link"`
		Miniprogram ReqSendWelcomeMsgExternalcontactMiniprogram `json:"miniprogram"`
		// Msgtype 附件类型，可选image、link、miniprogram或者video，必填
		Msgtype string                                `json:"msgtype"`
		Video   ReqSendWelcomeMsgExternalcontactVideo `json:"video"`
	} `json:"attachments,omitempty"` // 附件，最多可添加9个附件
	Text ReqSendWelcomeMsgExternalcontactText `json:"text"`
	// WelcomeCode 通过<a href="#15260/%E6%B7%BB%E5%8A%A0%E5%A4%96%E9%83%A8%E8%81%94%E7%B3%BB%E4%BA%BA%E4%BA%8B%E4%BB%B6" rel="nofollow">添加外部联系人事件</a>推送给企业的发送欢迎语的凭证，有效期为<strong>20秒</strong>，必填
	WelcomeCode string `json:"welcome_code"`
}

var _ urlValuer = ReqSendWelcomeMsgExternalcontact{}

func (x ReqSendWelcomeMsgExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespSendWelcomeMsgExternalcontact 发送新客户欢迎语响应
// 文档：https://developer.work.weixin.qq.com/document/path/92599#发送新客户欢迎语
type RespSendWelcomeMsgExternalcontact struct {
	CommonResp
}

var _ bodyer = RespSendWelcomeMsgExternalcontact{}

func (x RespSendWelcomeMsgExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecSendWelcomeMsgExternalcontact 发送新客户欢迎语
// 文档：https://developer.work.weixin.qq.com/document/path/92599#发送新客户欢迎语
func (c *ApiClient) ExecSendWelcomeMsgExternalcontact(req ReqSendWelcomeMsgExternalcontact) (RespSendWelcomeMsgExternalcontact, error) {
	var resp RespSendWelcomeMsgExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/send_welcome_msg", req, &resp, true)
	if err != nil {
		return RespSendWelcomeMsgExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSendWelcomeMsgExternalcontact{}, bizErr
	}
	return resp, nil
}
