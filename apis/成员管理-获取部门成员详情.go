package apis

import (
	"encoding/json"

	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqListUser 获取部门成员详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/90337#获取部门成员详情
type ReqListUser struct {
	// DepartmentID 获取的部门id，必填
	DepartmentID int `json:"department_id"`
	// FetchChild 1/0:是否递归获取子部门下面的成员
	FetchChild int `json:"fetch_child"`
}

var _ urlValuer = ReqListUser{}

func (x ReqListUser) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespListUser 获取部门成员详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/90337#获取部门成员详情
type RespListUser struct {
	CommonResp
	Userlist []struct {
		// Address 地址。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Address string `json:"address"`
		// Alias 别名；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Alias string `json:"alias"`
		// Avatar 头像url。 第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Avatar string `json:"avatar"`
		// BizMail 企业邮箱，代开发自建应用不返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		BizMail string `json:"biz_mail"`
		// Department 成员所属部门id列表，仅返回该应用有查看权限的部门id。对授权了“组织架构信息”的第三方应用，返回成员所属的全部部门id列表
		Department []int `json:"department"`
		// DirectLeader 直属上级UserID，返回在应用可见范围内的直属上级列表，最多有五个直属上级；第三方通讯录应用或者授权了“组织架构信息-应用可获取可见范围内成员组织架构信息-直属上级”权限的第三方应用可获取；对于非第三方创建的成员，第三方通讯录应用不可获取；上游企业不可获取下游企业成员该字段
		DirectLeader []string `json:"direct_leader"`
		// Email 邮箱，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Email       string `json:"email"`
		EnglishName string `json:"english_name"`
		Extattr     struct {
			Attrs []struct {
				// Name 成员名称；第三方不可获取，调用时返回userid以代替name；代开发自建应用需要管理员授权才返回；对于非第三方创建的成员，第三方通讯录应用也不可获取；未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示名字
				// Name 成员名称；第三方不可获取，调用时返回userid以代替name；代开发自建应用需要管理员授权才返回；对于非第三方创建的成员，第三方通讯录应用也不可获取；未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示名字
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text"`
				Type int `json:"type"`
				Web  struct {
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"web"`
			} `json:"attrs"`
		} `json:"extattr"` // 扩展属性，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		// ExternalPosition 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		ExternalPosition string `json:"external_position"`
		ExternalProfile  struct {
			ExternalAttr []struct {
				Miniprogram struct {
					Appid    string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Title    string `json:"title"`
				} `json:"miniprogram"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text"`
				Type int `json:"type"`
				Web  struct {
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"web"`
			} `json:"external_attr"`
			ExternalCorpName string `json:"external_corp_name"`
			WechatChannels   struct {
				Nickname string `json:"nickname"`
				// Status 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。<br/>已激活代表已激活企业微信或已关注微信插件（原企业号）。未激活代表既未激活企业微信又未关注微信插件（原企业号）。
				Status int `json:"status"`
			} `json:"wechat_channels"`
		} `json:"external_profile"` // 成员对外属性，字段详情见<a href="#13450" rel="nofollow">对外属性</a>；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		// Gender 性别。0表示未定义，1表示男性，2表示女性。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段。注:不可获取指返回值为0
		Gender string `json:"gender"`
		// IsLeaderInDept 表示在所在的部门内是否为部门负责人。0-否；1-是。是一个列表，数量必须与department一致。第三方通讯录应用或者授权了“组织架构信息-应用可获取企业的部门组织架构信息-部门负责人”权限的第三方应用可获取；对于非第三方创建的成员，第三方通讯录应用不可获取；上游企业不可获取下游企业成员该字段
		IsLeaderInDept []int `json:"is_leader_in_dept"`
		// MainDepartment 主部门，仅当应用对主部门有查看权限时返回。
		MainDepartment int `json:"main_department"`
		// Mobile 手机号码，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Mobile string `json:"mobile"`
		// Name 成员名称；第三方不可获取，调用时返回userid以代替name；代开发自建应用需要管理员授权才返回；对于非第三方创建的成员，第三方通讯录应用也不可获取；未返回名称的情况需要通过<a href="#17172" rel="nofollow">通讯录展示组件</a>来展示名字
		Name string `json:"name"`
		// OpenUserid 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
		OpenUserid string `json:"open_userid"`
		// Order 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。值范围是[0, 2^32)
		Order []int `json:"order"`
		// Position 职务信息；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Position string `json:"position"`
		// QrCode 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		QrCode string `json:"qr_code"`
		// Status 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。<br/>已激活代表已激活企业微信或已关注微信插件（原企业号）。未激活代表既未激活企业微信又未关注微信插件（原企业号）。
		Status int `json:"status"`
		// Telephone 座机。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		Telephone string `json:"telephone"`
		// ThumbAvatar 头像缩略图url。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取；上游企业不可获取下游企业成员该字段
		ThumbAvatar string `json:"thumb_avatar"`
		// Userid 成员UserID。对应管理端的帐号
		Userid string `json:"userid"`
	} `json:"userlist"` // 成员列表
}

var _ bodyer = RespListUser{}

func (x RespListUser) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecListUser 获取部门成员详情
// 文档：https://developer.work.weixin.qq.com/document/path/90337#获取部门成员详情
func (c *ApiClient) ExecListUser(req ReqListUser) (RespListUser, error) {
	var resp RespListUser
	err := c.executeWXApiGet("/cgi-bin/user/list", req, &resp, true)
	if err != nil {
		return RespListUser{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespListUser{}, bizErr
	}
	return resp, nil
}
