package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 修改生成的文件,以满足开发需求

// ReqAddContactWayExternalcontact 配置客户联系「联系我」方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#配置客户联系「联系我」方式
type (
	AddContactWayConclusions struct {
		Image       AddContactWayConclusionsImage       `json:"image"`
		Link        AddContactWayConclusionsLink        `json:"link"`
		Miniprogram AddContactWayConclusionsMiniprogram `json:"miniprogram"`
		Text        AddContactWayConclusionsText        `json:"text"`
	}

	AddContactWayConclusionsImage struct {
		MediaID string `json:"media_id"`
	}

	AddContactWayConclusionsLink struct {
		Desc   string `json:"desc"`
		Picurl string `json:"picurl"`
		Title  string `json:"title"`
		URL    string `json:"url"`
	}

	AddContactWayConclusionsMiniprogram struct {
		Appid      string `json:"appid"`
		Page       string `json:"page"`
		PicMediaID string `json:"pic_media_id"`
		Title      string `json:"title"`
	}

	AddContactWayConclusionsText struct {
		Content string `json:"content"`
	}
)

type ReqAddContactWayExternalcontact struct {
	// ChatExpiresIn 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
	ChatExpiresIn int                      `json:"chat_expires_in"`
	Conclusions   AddContactWayConclusions `json:"conclusions"` // 结束语，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅在is_temp为true时有效
	// ExpiresIn 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
	ExpiresIn int `json:"expires_in"`
	// IsTemp 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp"`
	// Party 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party"`
	// Remark 联系方式的备注信息，用于助记，不超过30个字符
	Remark string `json:"remark"`
	// Scene 场景，1-在小程序中联系，2-通过二维码联系，必填
	Scene int `json:"scene"`
	// SkipVerify 外部客户添加时是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“<a href="#13878">获取外部联系人详情</a>”时会返回该参数值，不超过30个字符
	State string `json:"state"`
	// Style 在小程序中联系时使用的控件样式，详见附表
	Style int `json:"style"`
	// Type 联系方式类型,1-单人, 2-多人，必填
	Type int `json:"type"`
	// Unionid 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	Unionid string `json:"unionid"`
	// User 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user"`
}

var _ bodyer = ReqAddContactWayExternalcontact{}

func (x ReqAddContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespAddContactWayExternalcontact 配置客户联系「联系我」方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#配置客户联系「联系我」方式
type RespAddContactWayExternalcontact struct {
	ConfigID string `json:"config_id"`
	CommonResp
	QrCode string `json:"qr_code"`
}

var _ bodyer = RespAddContactWayExternalcontact{}

func (x RespAddContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecAddContactWayExternalcontact 配置客户联系「联系我」方式
// 文档：https://developer.work.weixin.qq.com/document/path/92577#配置客户联系「联系我」方式
func (c *ApiClient) ExecAddContactWayExternalcontact(req ReqAddContactWayExternalcontact) (RespAddContactWayExternalcontact, error) {
	var resp RespAddContactWayExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/add_contact_way", req, &resp, true)
	if err != nil {
		return RespAddContactWayExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddContactWayExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqGetContactWayExternalcontact 获取企业已配置的「联系我」方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」方式
type ReqGetContactWayExternalcontact struct {
	// ConfigID 联系方式的配置id，必填
	ConfigID string `json:"config_id"`
}

var _ bodyer = ReqGetContactWayExternalcontact{}

func (x ReqGetContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetContactWayExternalcontact 获取企业已配置的「联系我」方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」方式
type RespGetContactWayExternalcontact struct {
	ContactWay struct {
		ChatExpiresIn int `json:"chat_expires_in"`
		Conclusions   struct {
			Image struct {
				PicURL string `json:"pic_url"`
			} `json:"image"`
			Link struct {
				Desc   string `json:"desc"`
				Picurl string `json:"picurl"`
				Title  string `json:"title"`
				URL    string `json:"url"`
			} `json:"link"`
			Miniprogram struct {
				Appid      string `json:"appid"`
				Page       string `json:"page"`
				PicMediaID string `json:"pic_media_id"`
				Title      string `json:"title"`
			} `json:"miniprogram"`
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"conclusions"`
		ConfigID   string   `json:"config_id"`
		ExpiresIn  int      `json:"expires_in"`
		IsTemp     bool     `json:"is_temp"`
		Party      []int    `json:"party"`
		QrCode     string   `json:"qr_code"`
		Remark     string   `json:"remark"`
		Scene      int      `json:"scene"`
		SkipVerify bool     `json:"skip_verify"`
		State      string   `json:"state"`
		Style      int      `json:"style"`
		Type       int      `json:"type"`
		Unionid    string   `json:"unionid"`
		User       []string `json:"user"`
	} `json:"contact_way"`
	CommonResp
}

var _ bodyer = RespGetContactWayExternalcontact{}

func (x RespGetContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetContactWayExternalcontact 获取企业已配置的「联系我」方式
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」方式
func (c *ApiClient) ExecGetContactWayExternalcontact(req ReqGetContactWayExternalcontact) (RespGetContactWayExternalcontact, error) {
	var resp RespGetContactWayExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/get_contact_way", req, &resp, true)
	if err != nil {
		return RespGetContactWayExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetContactWayExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqListContactWayExternalcontact 获取企业已配置的「联系我」列表请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」列表
type ReqListContactWayExternalcontact struct {
	// Cursor 分页查询使用的游标，为上次请求返回的 <code>next_cursor</code>
	Cursor string `json:"cursor"`
	// EndTime 「联系我」创建结束时间戳, 默认为当前时间
	EndTime int `json:"end_time"`
	// Limit 每次查询的分页大小，默认为100条，最多支持1000条
	Limit int `json:"limit"`
	// StartTime 「联系我」创建起始时间戳, 默认为90天前
	StartTime int `json:"start_time"`
}

var _ bodyer = ReqListContactWayExternalcontact{}

func (x ReqListContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespListContactWayExternalcontact 获取企业已配置的「联系我」列表响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」列表
type RespListContactWayExternalcontact struct {
	ContactWay []struct {
		ConfigID string `json:"config_id"`
	} `json:"contact_way"`
	CommonResp
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespListContactWayExternalcontact{}

func (x RespListContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListContactWayExternalcontact 获取企业已配置的「联系我」列表
// 文档：https://developer.work.weixin.qq.com/document/path/92577#获取企业已配置的「联系我」列表
func (c *ApiClient) ExecListContactWayExternalcontact(req ReqListContactWayExternalcontact) (RespListContactWayExternalcontact, error) {
	var resp RespListContactWayExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/list_contact_way", req, &resp, true)
	if err != nil {
		return RespListContactWayExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListContactWayExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqUpdateContactWayExternalcontact 更新企业已配置的「联系我」方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#更新企业已配置的「联系我」方式

type ReqUpdateContactWayExternalcontact struct {
	// ChatExpiresIn 临时会话有效期，以秒为单位，该参数仅在临时会话模式下有效
	ChatExpiresIn int                      `json:"chat_expires_in"`
	Conclusions   AddContactWayConclusions `json:"conclusions"` // 结束语，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅临时会话模式（is_temp为true）可设置
	// ConfigID 企业联系方式的配置id，必填
	ConfigID string `json:"config_id"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位，该参数仅在临时会话模式下有效
	ExpiresIn int `json:"expires_in"`
	// Party 使用该联系方式的部门列表，将覆盖原有部门列表，只在配置的type为2时有效
	Party []int `json:"party"`
	// Remark 联系方式的备注信息，不超过30个字符，将覆盖之前的备注
	Remark string `json:"remark"`
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify bool `json:"skip_verify"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“<a href="#13878">获取外部联系人详情</a>”时会返回该参数值
	State string `json:"state"`
	// Style 样式，只针对“在小程序中联系”的配置生效
	Style int `json:"style"`
	// Unionid 可进行临时会话的客户unionid，该参数仅在临时会话模式有效，如不指定则不进行限制
	Unionid string `json:"unionid"`
	// User 使用该联系方式的用户列表，将覆盖原有用户列表
	User []string `json:"user"`
}

var _ bodyer = ReqUpdateContactWayExternalcontact{}

func (x ReqUpdateContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateContactWayExternalcontact 更新企业已配置的「联系我」方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#更新企业已配置的「联系我」方式
type RespUpdateContactWayExternalcontact struct {
	CommonResp
}

var _ bodyer = RespUpdateContactWayExternalcontact{}

func (x RespUpdateContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateContactWayExternalcontact 更新企业已配置的「联系我」方式
// 文档：https://developer.work.weixin.qq.com/document/path/92577#更新企业已配置的「联系我」方式
func (c *ApiClient) ExecUpdateContactWayExternalcontact(req ReqUpdateContactWayExternalcontact) (RespUpdateContactWayExternalcontact, error) {
	var resp RespUpdateContactWayExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/update_contact_way", req, &resp, true)
	if err != nil {
		return RespUpdateContactWayExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateContactWayExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqDelContactWayExternalcontact 删除企业已配置的「联系我」方式请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#删除企业已配置的「联系我」方式
type ReqDelContactWayExternalcontact struct {
	// ConfigID 企业联系方式的配置id，必填
	ConfigID string `json:"config_id"`
}

var _ bodyer = ReqDelContactWayExternalcontact{}

func (x ReqDelContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespDelContactWayExternalcontact 删除企业已配置的「联系我」方式响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#删除企业已配置的「联系我」方式
type RespDelContactWayExternalcontact struct {
	CommonResp
}

var _ bodyer = RespDelContactWayExternalcontact{}

func (x RespDelContactWayExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecDelContactWayExternalcontact 删除企业已配置的「联系我」方式
// 文档：https://developer.work.weixin.qq.com/document/path/92577#删除企业已配置的「联系我」方式
func (c *ApiClient) ExecDelContactWayExternalcontact(req ReqDelContactWayExternalcontact) (RespDelContactWayExternalcontact, error) {
	var resp RespDelContactWayExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/del_contact_way", req, &resp, true)
	if err != nil {
		return RespDelContactWayExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespDelContactWayExternalcontact{}, bizErr
	}
	return resp, nil
}

// ReqCloseTempChatExternalcontact 结束临时会话请求
// 文档：https://developer.work.weixin.qq.com/document/path/92577#结束临时会话
type ReqCloseTempChatExternalcontact struct {
	// ExternalUserid 客户的外部联系人userid，必填
	ExternalUserid string `json:"external_userid"`
	// Userid 企业成员的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = ReqCloseTempChatExternalcontact{}

func (x ReqCloseTempChatExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCloseTempChatExternalcontact 结束临时会话响应
// 文档：https://developer.work.weixin.qq.com/document/path/92577#结束临时会话
type RespCloseTempChatExternalcontact struct {
	CommonResp
}

var _ bodyer = RespCloseTempChatExternalcontact{}

func (x RespCloseTempChatExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCloseTempChatExternalcontact 结束临时会话
// 文档：https://developer.work.weixin.qq.com/document/path/92577#结束临时会话
func (c *ApiClient) ExecCloseTempChatExternalcontact(req ReqCloseTempChatExternalcontact) (RespCloseTempChatExternalcontact, error) {
	var resp RespCloseTempChatExternalcontact
	err := c.executeWXApiPost("/cgi-bin/externalcontact/close_temp_chat", req, &resp, true)
	if err != nil {
		return RespCloseTempChatExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCloseTempChatExternalcontact{}, bizErr
	}
	return resp, nil
}
