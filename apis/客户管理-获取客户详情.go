package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetExternalcontact 获取客户详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/92265#获取客户详情
type ReqGetExternalcontact struct {
	// AccessToken 调用接口凭证，必填
	AccessToken string `json:"access_token"`
	// Cursor 上次请求返回的<code>next_cursor</code>
	Cursor string `json:"cursor"`
	// ExternalUserid 外部联系人的<code>userid</code>，注意不是企业成员的帐号，必填
	ExternalUserid string `json:"external_userid"`
}

var _ urlValuer = ReqGetExternalcontact{}

func (x ReqGetExternalcontact) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetExternalcontact 获取客户详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/92265#获取客户详情
type RespGetExternalcontact struct {
	CommonResp
	ExternalContact struct {
		// Avatar 外部联系人头像，代开发自建应用需要管理员授权才可以获取，第三方不可获取，上游企业不可获取下游企业客户该字段
		Avatar string `json:"avatar"`
		// CorpFullName 外部联系人所在企业的主体名称，<font data-color="#DC143C" color="#DC143C">仅当联系人类型是企业微信用户时有此字段</font>
		CorpFullName string `json:"corp_full_name"`
		// CorpName 外部联系人所在企业的简称，<font data-color="#DC143C" color="#DC143C">仅当联系人类型是企业微信用户时有此字段</font>
		CorpName        string `json:"corp_name"`
		ExternalProfile struct {
			ExternalAttr []struct {
				Miniprogram struct {
					Appid    string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Title    string `json:"title"`
				} `json:"miniprogram"`
				// Name 外部联系人的名称<sup>[注1]</sup>
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text"`
				// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
				// Type 该成员添加此外部联系人所打标签类型, 1-企业设置，2-用户自定义，3-规则组标签（<strong>仅系统应用返回</strong>）
				Type int `json:"type"`
				Web  struct {
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"web"`
			} `json:"external_attr"`
		} `json:"external_profile"` // 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，<font data-color="#DC143C" color="#DC143C">仅当联系人类型是企业微信用户时有此字段</font>，字段详情见<a href="#13450" rel="nofollow">对外属性</a>；
		// ExternalUserid 外部联系人的userid
		ExternalUserid string `json:"external_userid"`
		// Gender 外部联系人性别 0-未知 1-男性 2-女性。<strong>第三方不可获取</strong>，上游企业不可获取下游企业客户该字段，返回值为0，表示未定义
		Gender int `json:"gender"`
		// Name 外部联系人的名称<sup>[注1]</sup>
		Name string `json:"name"`
		// Position 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，<font data-color="#DC143C" color="#DC143C">仅当联系人类型是企业微信用户时有此字段</font>
		Position string `json:"position"`
		// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
		// Type 该成员添加此外部联系人所打标签类型, 1-企业设置，2-用户自定义，3-规则组标签（<strong>仅系统应用返回</strong>）
		Type int `json:"type"`
		// Unionid 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。<font data-color="#DC143C" color="#DC143C">仅当联系人类型是微信用户，且企业绑定了微信开发者ID有此字段</font>。<a href="#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%BE%AE%E4%BF%A1%E5%BC%80%E5%8F%91%E8%80%85id" rel="nofollow">查看绑定方法</a>。<strong>第三方不可获取</strong>，上游企业不可获取下游企业客户的unionid字段
		Unionid string `json:"unionid"`
	} `json:"external_contact"`
	FollowUser []struct {
		// AddWay 该成员添加此客户的来源，具体含义详见<a href="#13878/%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89" rel="nofollow">来源定义</a>
		AddWay int `json:"add_way"`
		// Createtime 该成员添加此外部联系人的时间
		Createtime int `json:"createtime"`
		// Description 该成员对此外部联系人的描述
		Description string `json:"description"`
		// OperUserid 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
		OperUserid string `json:"oper_userid"`
		// Remark 该成员对此外部联系人的备注
		Remark string `json:"remark"`
		// RemarkCorpName 该成员对此微信客户备注的企业名称（仅微信客户有该字段）
		RemarkCorpName string `json:"remark_corp_name"`
		// RemarkMobiles 该成员对此客户备注的手机号码，代开发自建应用需要管理员授权才可以获取，<strong>第三方不可获取</strong>，上游企业不可获取下游企业客户该字段
		RemarkMobiles []string `json:"remark_mobiles"`
		// State 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过<a href="#15645/%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F" rel="nofollow">创建「联系我」方式</a>指定
		State string `json:"state"`
		Tags  []struct {
			// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
			GroupName string `json:"group_name"`
			// TagID 该成员添加此外部联系人所打<strong>企业标签</strong>的id，<strong>用户自定义类型标签（type=2）不返回</strong>
			TagID string `json:"tag_id"`
			// TagName 该成员添加此外部联系人所打标签名称
			TagName string `json:"tag_name"`
			// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
			// Type 该成员添加此外部联系人所打标签类型, 1-企业设置，2-用户自定义，3-规则组标签（<strong>仅系统应用返回</strong>）
			Type int `json:"type"`
		} `json:"tags"`
		// Userid 添加了此外部联系人的企业成员userid
		Userid         string `json:"userid"`
		WechatChannels struct {
			// Nickname 视频号名称
			Nickname string `json:"nickname"`
		} `json:"wechat_channels"` // 该成员添加此客户的来源add_way为10时，对应的视频号信息
	} `json:"follow_user"`
	// NextCursor 分页的cursor，当跟进人多于500人时返回
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = RespGetExternalcontact{}

func (x RespGetExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetExternalcontact 获取客户详情
// 文档：https://developer.work.weixin.qq.com/document/path/92265#获取客户详情
func (c *ApiClient) ExecGetExternalcontact(req ReqGetExternalcontact) (RespGetExternalcontact, error) {
	var resp RespGetExternalcontact
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get", req, &resp, true)
	if err != nil {
		return RespGetExternalcontact{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetExternalcontact{}, bizErr
	}
	return resp, nil
}
