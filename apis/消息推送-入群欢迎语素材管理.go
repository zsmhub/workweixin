package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqAddGroupWelcomeTemplate 添加入群欢迎语素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/93438#添加入群欢迎语素材
type ReqAddGroupWelcomeTemplate struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid     int                                         `json:"agentid,omitempty"`
	File        ReqSendWelcomeMsgExternalcontactFile        `json:"file"`
	Image       ReqSendWelcomeMsgExternalcontactImage       `json:"image"`
	Link        ReqSendWelcomeMsgExternalcontactLink        `json:"link"`
	Miniprogram ReqSendWelcomeMsgExternalcontactMiniprogram `json:"miniprogram"`
	// Notify 是否通知成员将这条入群欢迎语应用到客户群中，0-不通知，1-通知， <strong>不填则通知</strong>
	Notify int                                         `json:"notify,omitempty"`
	Text   ReqSendWelcomeMsgExternalcontactText        `json:"text"`
	Video  ReqSendWelcomeMsgExternalcontactMiniprogram `json:"video"`
}

var _ urlValuer = ReqAddGroupWelcomeTemplate{}

func (x ReqAddGroupWelcomeTemplate) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespAddGroupWelcomeTemplate 添加入群欢迎语素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/93438#添加入群欢迎语素材
type RespAddGroupWelcomeTemplate struct {
	CommonResp
	// TemplateID 欢迎语素材id
	TemplateID string `json:"template_id"`
}

var _ bodyer = RespAddGroupWelcomeTemplate{}

func (x RespAddGroupWelcomeTemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddGroupWelcomeTemplate 添加入群欢迎语素材
// 文档：https://developer.work.weixin.qq.com/document/path/93438#添加入群欢迎语素材
func (c *ApiClient) ExecAddGroupWelcomeTemplate(req ReqAddGroupWelcomeTemplate) (RespAddGroupWelcomeTemplate, error) {
	var resp RespAddGroupWelcomeTemplate
	err := c.executeWXApiGet("/cgi-bin/externalcontact/group_welcome_template/add", req, &resp, true)
	if err != nil {
		return RespAddGroupWelcomeTemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddGroupWelcomeTemplate{}, bizErr
	}
	return resp, nil
}

// ReqEditGroupWelcomeTemplate 编辑入群欢迎语素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/93438#编辑入群欢迎语素材
type ReqEditGroupWelcomeTemplate struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid     int                                         `json:"agentid,omitempty"`
	File        ReqSendWelcomeMsgExternalcontactFile        `json:"file"`
	Image       ReqSendWelcomeMsgExternalcontactFile        `json:"image"`
	Link        ReqSendWelcomeMsgExternalcontactLink        `json:"link"`
	Miniprogram ReqSendWelcomeMsgExternalcontactMiniprogram `json:"miniprogram"`
	// TemplateID 欢迎语素材id，必填
	TemplateID string                                `json:"template_id"`
	Text       ReqSendWelcomeMsgExternalcontactText  `json:"text"`
	Video      ReqSendWelcomeMsgExternalcontactVideo `json:"video"`
}

var _ urlValuer = ReqEditGroupWelcomeTemplate{}

func (x ReqEditGroupWelcomeTemplate) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespEditGroupWelcomeTemplate 编辑入群欢迎语素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/93438#编辑入群欢迎语素材
type RespEditGroupWelcomeTemplate struct {
	CommonResp
}

var _ bodyer = RespEditGroupWelcomeTemplate{}

func (x RespEditGroupWelcomeTemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecEditGroupWelcomeTemplate 编辑入群欢迎语素材
// 文档：https://developer.work.weixin.qq.com/document/path/93438#编辑入群欢迎语素材
func (c *ApiClient) ExecEditGroupWelcomeTemplate(req ReqEditGroupWelcomeTemplate) (RespEditGroupWelcomeTemplate, error) {
	var resp RespEditGroupWelcomeTemplate
	err := c.executeWXApiGet("/cgi-bin/externalcontact/group_welcome_template/edit", req, &resp, true)
	if err != nil {
		return RespEditGroupWelcomeTemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespEditGroupWelcomeTemplate{}, bizErr
	}
	return resp, nil
}

// ReqGetGroupWelcomeTemplate 获取入群欢迎语素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/93438#获取入群欢迎语素材
type ReqGetGroupWelcomeTemplate struct {
	// TemplateID 群欢迎语的素材id，必填
	TemplateID string `json:"template_id"`
}

var _ urlValuer = ReqGetGroupWelcomeTemplate{}

func (x ReqGetGroupWelcomeTemplate) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetGroupWelcomeTemplate 获取入群欢迎语素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/93438#获取入群欢迎语素材
type RespGetGroupWelcomeTemplate struct {
	CommonResp
	File        ReqSendWelcomeMsgExternalcontactFile        `json:"file"`
	Image       ReqSendWelcomeMsgExternalcontactImage       `json:"image"`
	Link        ReqSendWelcomeMsgExternalcontactLink        `json:"link"`
	Miniprogram ReqSendWelcomeMsgExternalcontactMiniprogram `json:"miniprogram"`
	Text        ReqSendWelcomeMsgExternalcontactText        `json:"text"`
	Video       ReqSendWelcomeMsgExternalcontactVideo       `json:"video"`
}

var _ bodyer = RespGetGroupWelcomeTemplate{}

func (x RespGetGroupWelcomeTemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetGroupWelcomeTemplate 获取入群欢迎语素材
// 文档：https://developer.work.weixin.qq.com/document/path/93438#获取入群欢迎语素材
func (c *ApiClient) ExecGetGroupWelcomeTemplate(req ReqGetGroupWelcomeTemplate) (RespGetGroupWelcomeTemplate, error) {
	var resp RespGetGroupWelcomeTemplate
	err := c.executeWXApiGet("/cgi-bin/externalcontact/group_welcome_template/get", req, &resp, true)
	if err != nil {
		return RespGetGroupWelcomeTemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetGroupWelcomeTemplate{}, bizErr
	}
	return resp, nil
}

// ReqDelGroupWelcomeTemplate 删除入群欢迎语素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/93438#删除入群欢迎语素材
type ReqDelGroupWelcomeTemplate struct {
	// Agentid 授权方安装的应用agentid。<strong>仅旧的第三方多应用套件需要填此参数</strong>
	Agentid int `json:"agentid,omitempty"`
	// TemplateID 群欢迎语的素材id，必填
	TemplateID string `json:"template_id"`
}

var _ urlValuer = ReqDelGroupWelcomeTemplate{}

func (x ReqDelGroupWelcomeTemplate) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespDelGroupWelcomeTemplate 删除入群欢迎语素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/93438#删除入群欢迎语素材
type RespDelGroupWelcomeTemplate struct {
	CommonResp
}

var _ bodyer = RespDelGroupWelcomeTemplate{}

func (x RespDelGroupWelcomeTemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDelGroupWelcomeTemplate 删除入群欢迎语素材
// 文档：https://developer.work.weixin.qq.com/document/path/93438#删除入群欢迎语素材
func (c *ApiClient) ExecDelGroupWelcomeTemplate(req ReqDelGroupWelcomeTemplate) (RespDelGroupWelcomeTemplate, error) {
	var resp RespDelGroupWelcomeTemplate
	err := c.executeWXApiGet("/cgi-bin/externalcontact/group_welcome_template/del", req, &resp, true)
	if err != nil {
		return RespDelGroupWelcomeTemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDelGroupWelcomeTemplate{}, bizErr
	}
	return resp, nil
}
